package dex

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/uatu/internal/contracts/uniswap"
)

type Client struct {
	client *ethclient.Client
	cow    *cowClient
}

type Permit2Allowance struct {
	Amount     *big.Int
	Expiration *big.Int
	Nonce      *big.Int
}

var erc20ABI = func() abi.ABI {
	parsed, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"approve","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"spender","type":"address"}],"name":"allowance","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`))
	if err != nil {
		panic(err)
	}
	return parsed
}()

func Provider(rpcUrl string) (*Client, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, fmt.Errorf("could not create eth client: %w", err)
	}
	return &Client{client: client, cow: newCowClient()}, nil
}

func (c *Client) getERC20Allowance(
	ctx context.Context,
	token, owner, spender common.Address,
) (*big.Int, error) {
	contract := bind.NewBoundContract(token, erc20ABI, c.client, nil, nil)
	var out []interface{}
	if err := contract.Call(&bind.CallOpts{Context: ctx}, &out, "allowance", owner, spender); err != nil {
		return nil, fmt.Errorf("could not get erc20 allowance: %w", err)
	}
	allowance, ok := out[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("unexpected erc20 allowance type %T", out[0])
	}
	return allowance, nil
}

func (c *Client) GetTokenBalance(
	ctx context.Context,
	token, account common.Address,
) (*big.Int, error) {
	contract := bind.NewBoundContract(token, erc20ABI, c.client, nil, nil)
	var out []interface{}
	if err := contract.Call(&bind.CallOpts{Context: ctx}, &out, "balanceOf", account); err != nil {
		return nil, fmt.Errorf("could not get token balance: %w", err)
	}
	balance, ok := out[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("unexpected token balance type %T", out[0])
	}
	return balance, nil
}

func encodeERC20Token(amount *big.Int, spender common.Address) ([]byte, error) {
	encoded, err := erc20ABI.Pack("approve", spender, amount)
	if err != nil {
		return nil, fmt.Errorf("could not encode approve calldata: %w", err)
	}
	return encoded, nil
}

func (c *Client) getPermit2TokenAllowance(
	ctx context.Context,
	permit2Address, wallet, token, spender common.Address,
) (*Permit2Allowance, error) {
	contract, err := uniswap.NewPermit2Caller(permit2Address, c.client)
	if err != nil {
		return nil, fmt.Errorf("could not bind permit2: %w", err)
	}
	allowance, err := contract.Allowance(
		&bind.CallOpts{Context: ctx},
		wallet,
		token,
		spender,
	)
	if err != nil {
		return nil, fmt.Errorf("could not get permit2 allowance: %w", err)
	}
	return &Permit2Allowance{
		Amount:     allowance.Amount,
		Expiration: allowance.Expiration,
		Nonce:      allowance.Nonce,
	}, nil
}

func encodePermit2Approval(token, spender common.Address, amount, expiration *big.Int) ([]byte, error) {
	permit2ABI, err := uniswap.Permit2MetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("could not parse permit2 abi: %w", err)
	}
	calldata, err := permit2ABI.Pack("approve", token, spender, amount, expiration)
	if err != nil {
		return nil, fmt.Errorf("could not encode permit2 approve calldata: %w", err)
	}
	return calldata, nil
}

func ScaleDecimals(amt *big.Int, from, to uint8) *big.Int {
	switch {
	case from == to:
		return new(big.Int).Set(amt)
	case from > to:
		divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(from-to)), nil)
		return new(big.Int).Div(amt, divisor)
	default:
		multiplier := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(to-from)), nil)
		return new(big.Int).Mul(amt, multiplier)
	}
}
