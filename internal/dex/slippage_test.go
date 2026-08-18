package dex

import (
	"math/big"
	"testing"
)

func TestApplySlippage(t *testing.T) {
	tests := []struct {
		name        string
		amount      string
		slippageBps uint
		want        string
	}{
		{"zero tolerance returns the quote", "1000000", 0, "1000000"},
		{"half a percent", "1000000", 50, "995000"},
		{"one percent", "1000000", 100, "990000"},
		{"the configured ceiling", "1000000", 5000, "500000"},
		{"rounds down rather than up", "1999", 50, "1989"},
		{"whole amount tolerated", "1000000", 10000, "0"},
		{"beyond the whole amount is floored at zero", "1000000", 20000, "0"},
		{"survives amounts wider than uint64", "123456789012345678901234567890", 250, "120370369287037036928703703692"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			amount, ok := new(big.Int).SetString(tt.amount, 10)
			if !ok {
				t.Fatalf("could not parse amount %q", tt.amount)
			}
			got := applySlippage(amount, tt.slippageBps)
			if got.String() != tt.want {
				t.Errorf("applySlippage(%s, %d) = %s, want %s",
					tt.amount, tt.slippageBps, got, tt.want)
			}
			if amount.String() != tt.amount {
				t.Errorf("applySlippage mutated its input: got %s, want %s", amount, tt.amount)
			}
		})
	}
}

func TestApplySlippageNil(t *testing.T) {
	if got := applySlippage(nil, 50); got != nil {
		t.Errorf("applySlippage(nil, 50) = %v, want nil", got)
	}
}

// CoW orders need headroom below the quote for solvers to recover gas, so a
// tighter request is raised to the floor instead of being honoured.
func TestCowSlippage(t *testing.T) {
	tests := []struct {
		requested uint
		want      uint
	}{
		{0, cowMinSlippageBps},
		{10, cowMinSlippageBps},
		{cowMinSlippageBps, cowMinSlippageBps},
		{100, 100},
		{5000, 5000},
	}
	for _, tt := range tests {
		if got := cowSlippage(tt.requested); got != tt.want {
			t.Errorf("cowSlippage(%d) = %d, want %d", tt.requested, got, tt.want)
		}
	}
}
