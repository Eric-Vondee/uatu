package dex

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/uatu"
)

// CowSwap settles off-chain: solvers batch orders and submit them to the
// GPv2Settlement contract, so there is no router to call and no pool to price
// against. An order is a signed message rather than a transaction. This
// integration uses the protocol's "presign" scheme, where the order is signed
// by calling setPreSignature on the settlement contract, which keeps the whole
// flow expressible as the transactions the rest of the codebase already emits:
//
//	1. approve the vault relayer to pull the sell token (when needed)
//	2. setPreSignature(orderUid, true) on the settlement contract
//
// The order itself is registered with the orderbook out of band, by
// IDexResponse.RegisterOrder, once a route has been selected.

const (
	// CowSlug is the dex slug cowswap is seeded under.
	CowSlug = "cowswap"

	cowAPIBaseURL = "https://api.cow.fi"

	// GPv2Settlement and GPv2VaultRelayer are deployed at the same address on
	// every chain CoW supports, so these serve as fallbacks for dex rows that
	// do not carry a settlement address.
	cowSettlementAddress   = "0x9008D19f58AAbD9eD0D60971565AA8510560ab41"
	cowVaultRelayerAddress = "0xC92E8bdf79f0507f65a392b0ab4667716BFE0110"

	// cowAppData is the app-data document hashed into an order's appData field.
	// "{}" is the documented value for integrations that attach no metadata.
	cowAppData = "{}"

	// cowMinSlippageBps is the floor on how far below the quoted output the
	// signed limit price may sit. Orders are posted with a zero feeAmount
	// because the orderbook rejects anything else, which means solvers recover
	// the network cost out of the difference between the limit price and what
	// they actually fill at. Signing the quote exactly would leave nothing to
	// pay for gas and the order would never be picked up, so a caller asking
	// for tighter slippage than this gets the floor instead.
	cowMinSlippageBps = 50

	// cowOrderValidity is how long a posted order stays fillable.
	cowOrderValidity = swapDeadline

	cowHTTPTimeout      = 15 * time.Second
	cowMaxResponseBytes = 1 << 20
)

// cowNetworks maps a chain id onto the path segment CoW's orderbook API is
// served under. A chain that is absent is not supported by the protocol.
var cowNetworks = map[uint]string{
	1:     "mainnet",
	56:    "bnb",
	100:   "xdai",
	137:   "polygon",
	8453:  "base",
	9745:  "plasma",
	42161: "arbitrum_one",
	43114: "avalanche",
	57073: "ink",
	59144: "linea",
}

// EIP-712 constants from GPv2Order and GPv2Signing. Deriving the hashes from
// the type strings rather than hard-coding the digests keeps them checkable
// against the contracts.
var (
	cowOrderTypeHash = crypto.Keccak256Hash([]byte(
		"Order(address sellToken,address buyToken,address receiver," +
			"uint256 sellAmount,uint256 buyAmount,uint32 validTo,bytes32 appData," +
			"uint256 feeAmount,string kind,bool partiallyFillable," +
			"string sellTokenBalance,string buyTokenBalance)"))
	cowDomainTypeHash = crypto.Keccak256Hash([]byte(
		"EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"))
	cowDomainName    = crypto.Keccak256Hash([]byte("Gnosis Protocol"))
	cowDomainVersion = crypto.Keccak256Hash([]byte("v2"))
	cowAppDataHash   = crypto.Keccak256Hash([]byte(cowAppData))
)

const (
	cowKindSell         = "sell"
	cowBalanceERC20     = "erc20"
	cowSigningPresign   = "presign"
	cowPresignSignature = "0x"

	// cowUIDLength is the packed width of an order uid: digest, owner, validTo.
	cowUIDLength = 32 + 20 + 4
)

// CowOrder is the GPv2 order struct. Its EIP-712 digest, together with the
// owner and expiry, forms the order uid the settlement contract pre-signs.
type CowOrder struct {
	SellToken         common.Address
	BuyToken          common.Address
	Receiver          common.Address
	SellAmount        *big.Int
	BuyAmount         *big.Int
	ValidTo           uint32
	AppData           common.Hash
	FeeAmount         *big.Int
	Kind              string
	PartiallyFillable bool
	SellTokenBalance  string
	BuyTokenBalance   string
}

// cowDomainSeparator is the EIP-712 domain separator of the settlement contract
// on the given chain.
func cowDomainSeparator(chainID uint, settlement common.Address) common.Hash {
	return crypto.Keccak256Hash(
		cowDomainTypeHash.Bytes(),
		cowDomainName.Bytes(),
		cowDomainVersion.Bytes(),
		common.LeftPadBytes(new(big.Int).SetUint64(uint64(chainID)).Bytes(), 32),
		common.LeftPadBytes(settlement.Bytes(), 32),
	)
}

func cowBoolWord(b bool) []byte {
	word := make([]byte, 32)
	if b {
		word[31] = 1
	}
	return word
}

// Digest returns the EIP-712 signing hash of the order, which is what the
// settlement contract derives an order uid from.
func (o CowOrder) Digest(chainID uint, settlement common.Address) common.Hash {
	structHash := crypto.Keccak256(
		cowOrderTypeHash.Bytes(),
		common.LeftPadBytes(o.SellToken.Bytes(), 32),
		common.LeftPadBytes(o.BuyToken.Bytes(), 32),
		common.LeftPadBytes(o.Receiver.Bytes(), 32),
		common.LeftPadBytes(o.SellAmount.Bytes(), 32),
		common.LeftPadBytes(o.BuyAmount.Bytes(), 32),
		common.LeftPadBytes(new(big.Int).SetUint64(uint64(o.ValidTo)).Bytes(), 32),
		o.AppData.Bytes(),
		common.LeftPadBytes(o.FeeAmount.Bytes(), 32),
		crypto.Keccak256([]byte(o.Kind)),
		cowBoolWord(o.PartiallyFillable),
		crypto.Keccak256([]byte(o.SellTokenBalance)),
		crypto.Keccak256([]byte(o.BuyTokenBalance)),
	)
	separator := cowDomainSeparator(chainID, settlement)
	return crypto.Keccak256Hash([]byte{0x19, 0x01}, separator.Bytes(), structHash)
}

// CowOrderUID packs the order identifier the protocol addresses orders by: the
// EIP-712 digest, the owner, and the expiry, concatenated without padding.
func CowOrderUID(digest common.Hash, owner common.Address, validTo uint32) []byte {
	uid := make([]byte, 0, cowUIDLength)
	uid = append(uid, digest.Bytes()...)
	uid = append(uid, owner.Bytes()...)
	return binary.BigEndian.AppendUint32(uid, validTo)
}

var cowSettlementABI = func() abi.ABI {
	parsed, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"bytes","name":"orderUid","type":"bytes"},{"internalType":"bool","name":"signed","type":"bool"}],"name":"setPreSignature","outputs":[],"stateMutability":"nonpayable","type":"function"}]`))
	if err != nil {
		panic(err)
	}
	return parsed
}()

// EncodeSetPreSignature encodes the settlement call that signs an order
// on-chain, standing in for an EIP-712 signature the owner never produces.
func EncodeSetPreSignature(orderUID []byte, signed bool) ([]byte, error) {
	if len(orderUID) != cowUIDLength {
		return nil, fmt.Errorf("order uid must be %d bytes, got %d", cowUIDLength, len(orderUID))
	}
	calldata, err := cowSettlementABI.Pack("setPreSignature", orderUID, signed)
	if err != nil {
		return nil, fmt.Errorf("could not encode setPreSignature calldata: %w", err)
	}
	return calldata, nil
}

// cowSettlement resolves the settlement contract for a dex row, falling back to
// the canonical deployment when the row carries no address.
func cowSettlement(d uatu.Dex) common.Address {
	if d.SettlementAddress != "" {
		return uatu.FormatEvmAddress(d.SettlementAddress)
	}
	return common.HexToAddress(cowSettlementAddress)
}

type cowClient struct {
	http    *http.Client
	baseURL string
}

func newCowClient() *cowClient {
	return &cowClient{
		http:    &http.Client{Timeout: cowHTTPTimeout},
		baseURL: cowAPIBaseURL,
	}
}

type cowAPIError struct {
	ErrorType   string `json:"errorType"`
	Description string `json:"description"`
}

func (e cowAPIError) Error() string {
	return fmt.Sprintf("cow api rejected the request: %s: %s", e.ErrorType, e.Description)
}

func (c *cowClient) post(ctx context.Context, network, path string, body, out any) error {
	payload, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("could not encode cow request: %w", err)
	}
	url := fmt.Sprintf("%s/%s/api/v1/%s", c.baseURL, network, path)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("could not build cow request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("could not reach cow api: %w", err)
	}
	defer func() { _ = res.Body.Close() }()

	data, err := io.ReadAll(io.LimitReader(res.Body, cowMaxResponseBytes))
	if err != nil {
		return fmt.Errorf("could not read cow response: %w", err)
	}
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		var apiErr cowAPIError
		if json.Unmarshal(data, &apiErr) == nil && apiErr.ErrorType != "" {
			return apiErr
		}
		return fmt.Errorf("cow api returned %s: %s", res.Status, strings.TrimSpace(string(data)))
	}
	if err := json.Unmarshal(data, out); err != nil {
		return fmt.Errorf("could not decode cow response: %w", err)
	}
	return nil
}

type cowQuoteRequest struct {
	SellToken           string `json:"sellToken"`
	BuyToken            string `json:"buyToken"`
	Receiver            string `json:"receiver"`
	From                string `json:"from"`
	Kind                string `json:"kind"`
	SellAmountBeforeFee string `json:"sellAmountBeforeFee"`
	ValidFor            uint32 `json:"validFor"`
	AppData             string `json:"appData"`
	AppDataHash         string `json:"appDataHash,omitempty"`
	SigningScheme       string `json:"signingScheme"`
	PriceQuality        string `json:"priceQuality,omitempty"`
}

type cowQuotedOrder struct {
	SellToken   string `json:"sellToken"`
	BuyToken    string `json:"buyToken"`
	SellAmount  string `json:"sellAmount"`
	BuyAmount   string `json:"buyAmount"`
	FeeAmount   string `json:"feeAmount"`
	ValidTo     uint32 `json:"validTo"`
	AppDataHash string `json:"appDataHash"`
	Kind        string `json:"kind"`
}

type cowQuoteResponse struct {
	Quote      cowQuotedOrder `json:"quote"`
	From       string         `json:"from"`
	ID         *int64         `json:"id"`
	Expiration time.Time      `json:"expiration"`
}

// Quote prices a sell order. The returned sellAmount is net of the network
// cost, which the response reports separately as feeAmount.
func (c *cowClient) Quote(ctx context.Context, network string, req cowQuoteRequest) (*cowQuoteResponse, error) {
	var res cowQuoteResponse
	if err := c.post(ctx, network, "quote", req, &res); err != nil {
		return nil, err
	}
	if res.Quote.AppDataHash != "" &&
		!strings.EqualFold(res.Quote.AppDataHash, cowAppDataHash.Hex()) {
		return nil, fmt.Errorf(
			"cow quote returned app data %s, expected %s",
			res.Quote.AppDataHash, cowAppDataHash.Hex(),
		)
	}
	if res.From != "" && !strings.EqualFold(res.From, req.From) {
		return nil, fmt.Errorf("cow quote returned from %s, expected %s", res.From, req.From)
	}
	if res.Quote.SellToken != "" && !strings.EqualFold(res.Quote.SellToken, req.SellToken) {
		return nil, fmt.Errorf("cow quote returned sell token %s, expected %s", res.Quote.SellToken, req.SellToken)
	}
	if res.Quote.BuyToken != "" && !strings.EqualFold(res.Quote.BuyToken, req.BuyToken) {
		return nil, fmt.Errorf("cow quote returned buy token %s, expected %s", res.Quote.BuyToken, req.BuyToken)
	}
	if res.Quote.Kind != "" && !strings.EqualFold(res.Quote.Kind, req.Kind) {
		return nil, fmt.Errorf("cow quote returned kind %s, expected %s", res.Quote.Kind, req.Kind)
	}
	return &res, nil
}

type cowOrderCreation struct {
	SellToken         string `json:"sellToken"`
	BuyToken          string `json:"buyToken"`
	Receiver          string `json:"receiver"`
	SellAmount        string `json:"sellAmount"`
	BuyAmount         string `json:"buyAmount"`
	ValidTo           uint32 `json:"validTo"`
	FeeAmount         string `json:"feeAmount"`
	Kind              string `json:"kind"`
	PartiallyFillable bool   `json:"partiallyFillable"`
	SellTokenBalance  string `json:"sellTokenBalance"`
	BuyTokenBalance   string `json:"buyTokenBalance"`
	AppData           string `json:"appData"`
	AppDataHash       string `json:"appDataHash"`
	SigningScheme     string `json:"signingScheme"`
	Signature         string `json:"signature"`
	From              string `json:"from"`
	QuoteID           *int64 `json:"quoteId,omitempty"`
}

// PlaceOrder registers the order with the orderbook so solvers can see it. The
// order is inert until the owner pre-signs it on-chain, which is why the
// orderbook accepts it before the sell token has been approved.
//
// It verifies that the uid the orderbook assigns matches the one derived
// locally: the pre-signature calldata is built against the local uid, so a
// mismatch would sign an order that does not exist.
func (c *cowClient) PlaceOrder(
	ctx context.Context,
	network string,
	order CowOrder,
	owner common.Address,
	quoteID *int64,
	expectedUID []byte,
) error {
	body := cowOrderCreation{
		SellToken:         order.SellToken.Hex(),
		BuyToken:          order.BuyToken.Hex(),
		Receiver:          order.Receiver.Hex(),
		SellAmount:        order.SellAmount.String(),
		BuyAmount:         order.BuyAmount.String(),
		ValidTo:           order.ValidTo,
		FeeAmount:         order.FeeAmount.String(),
		Kind:              order.Kind,
		PartiallyFillable: order.PartiallyFillable,
		SellTokenBalance:  order.SellTokenBalance,
		BuyTokenBalance:   order.BuyTokenBalance,
		AppData:           cowAppData,
		AppDataHash:       cowAppDataHash.Hex(),
		SigningScheme:     cowSigningPresign,
		Signature:         cowPresignSignature,
		From:              owner.Hex(),
		QuoteID:           quoteID,
	}

	var uid string
	if err := c.post(ctx, network, "orders", body, &uid); err != nil {
		return err
	}
	if !strings.EqualFold(uid, uatu.HexBytes(expectedUID)) {
		return fmt.Errorf(
			"cow orderbook returned uid %s, expected %s",
			uid, uatu.HexBytes(expectedUID),
		)
	}
	return nil
}

func cowSlippage(slippageBps uint) uint {
	if slippageBps < cowMinSlippageBps {
		return cowMinSlippageBps
	}
	return slippageBps
}

func cowAmount(field, value string) (*big.Int, error) {
	amount, ok := new(big.Int).SetString(value, 10)
	if !ok || amount.Sign() < 0 {
		return nil, fmt.Errorf("could not parse cow quote %s %q", field, value)
	}
	return amount, nil
}

// SwapCow prices a swap against CoW's solver competition and encodes the
// transactions that authorise it. The order is posted through RegisterOrder only
// after this response wins route selection; otherwise every candidate quote
// would leave an unused order in CoW's orderbook.
func (c *Client) CowSwap(ctx context.Context, d uatu.IDexRequest) (*uatu.IDexResponse, error) {
	network, ok := cowNetworks[d.ChainId]
	if !ok {
		return nil, fmt.Errorf("cowswap is not deployed on chain %d", d.ChainId)
	}
	// Native sells go through CoW's separate ethflow contract, which this
	// integration does not cover.
	if d.TokenIn == (common.Address{}) || d.TokenOut == (common.Address{}) {
		return nil, fmt.Errorf("cowswap only swaps erc20 tokens")
	}
	if d.AmountIn == nil || d.AmountIn.Sign() <= 0 {
		return nil, fmt.Errorf("cowswap amount in must be positive")
	}

	settlement := cowSettlement(d.Dex)
	relayer := common.HexToAddress(cowVaultRelayerAddress)
	cow := c.cow
	if cow == nil {
		cow = newCowClient()
	}

	quote, err := cow.Quote(ctx, network, cowQuoteRequest{
		SellToken:           d.TokenIn.Hex(),
		BuyToken:            d.TokenOut.Hex(),
		Receiver:            d.WalletAddress.Hex(),
		From:                d.WalletAddress.Hex(),
		Kind:                cowKindSell,
		SellAmountBeforeFee: d.AmountIn.String(),
		ValidFor:            uint32(cowOrderValidity.Seconds()),
		AppData:             cowAppData,
		AppDataHash:         cowAppDataHash.Hex(),
		SigningScheme:       cowSigningPresign,
		PriceQuality:        "optimal",
	})
	if err != nil {
		return nil, err
	}
	if !quote.Expiration.IsZero() && !time.Now().Before(quote.Expiration) {
		return nil, fmt.Errorf("cow quote expired at %s", quote.Expiration.UTC().Format(time.RFC3339Nano))
	}
	if quote.Expiration.IsZero() {
		return nil, fmt.Errorf("cow quote did not include an expiration")
	}
	if quote.Quote.ValidTo <= uint32(time.Now().Unix()) {
		return nil, fmt.Errorf("cow quote order is already expired")
	}

	sellAmount, err := cowAmount("sellAmount", quote.Quote.SellAmount)
	if err != nil {
		return nil, err
	}
	buyAmount, err := cowAmount("buyAmount", quote.Quote.BuyAmount)
	if err != nil {
		return nil, err
	}
	if sellAmount.Sign() == 0 || buyAmount.Sign() == 0 {
		return nil, fmt.Errorf("cow quote returned a zero trade amount")
	}
	feeAmount, err := cowAmount("feeAmount", quote.Quote.FeeAmount)
	if err != nil {
		return nil, err
	}

	// The quote splits the requested amount into what is actually sold and the
	// network cost. The order has to be signed for the total, because the
	// settlement contract takes the cost out of the sell amount, and its own
	// feeAmount has to be zero: the orderbook rejects orders that set one.
	signedSellAmount := new(big.Int).Add(sellAmount, feeAmount)
	if signedSellAmount.Cmp(d.AmountIn) > 0 {
		return nil, fmt.Errorf("cow quote sell amount %s exceeds requested amount %s", signedSellAmount, d.AmountIn)
	}
	minimumBuyAmount := applySlippage(buyAmount, cowSlippage(d.SlippageBps))
	if minimumBuyAmount.Sign() == 0 {
		return nil, fmt.Errorf("cow quote minimum buy amount is zero")
	}

	order := CowOrder{
		SellToken:         d.TokenIn,
		BuyToken:          d.TokenOut,
		Receiver:          d.WalletAddress,
		SellAmount:        signedSellAmount,
		BuyAmount:         minimumBuyAmount,
		ValidTo:           quote.Quote.ValidTo,
		AppData:           cowAppDataHash,
		FeeAmount:         big.NewInt(0),
		Kind:              cowKindSell,
		PartiallyFillable: false,
		SellTokenBalance:  cowBalanceERC20,
		BuyTokenBalance:   cowBalanceERC20,
	}

	uid := CowOrderUID(order.Digest(d.ChainId, settlement), d.WalletAddress, order.ValidTo)
	preSignCalldata, err := EncodeSetPreSignature(uid, true)
	if err != nil {
		return nil, err
	}

	// CoW pulls the sell token through its vault relayer, not permit2.
	var erc20Approval []byte
	allowance, err := c.getERC20Allowance(ctx, d.TokenIn, d.WalletAddress, relayer)
	if err != nil {
		return nil, err
	}
	if allowance.Cmp(signedSellAmount) < 0 {
		erc20Approval, err = encodeERC20Token(signedSellAmount, relayer)
		if err != nil {
			return nil, err
		}
	}

	quoteID := quote.ID
	return &uatu.IDexResponse{
		AmountIn:             signedSellAmount,
		AmountOut:            buyAmount,
		AmountOutMinimum:     minimumBuyAmount,
		EncodedData:          preSignCalldata,
		EncodedERC20Approval: erc20Approval,
		Dex:                  d.Dex,
		RouterAddress:        settlement,
		RegisterOrder: func(ctx context.Context) error {
			return cow.PlaceOrder(ctx, network, order, d.WalletAddress, quoteID, uid)
		},
		Route: routeWithFee(DexRoutes["cowswap"], feeAmount),
	}, nil
}
