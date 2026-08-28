package uatu

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func FormatEvmAddress(address string) common.Address {
	return common.HexToAddress(address)
}

func ParseEVMAddress(address string) (common.Address, error) {
	address = strings.TrimSpace(address)
	if !common.IsHexAddress(address) {
		return common.Address{}, fmt.Errorf("invalid EVM address")
	}
	return common.HexToAddress(address), nil
}

func HexBytes(b []byte) string {
	return "0x" + hex.EncodeToString(b)
}
