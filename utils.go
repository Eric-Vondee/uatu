package uatu

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
)

func FormatEvmAddress(address string) common.Address {
	return common.HexToAddress(address)
}

func HexBytes(b []byte) string {
	return "0x" + hex.EncodeToString(b)
}
