package feeds

import (
	"math/big"
	"testing"
)

func TestFormatAnswer(t *testing.T) {
	tests := []struct {
		name      string
		answer    string
		decimals  uint8
		wantPrice string
	}{
		{name: "eight decimals", answer: "123456789", decimals: 8, wantPrice: "1.234567890000000000"},
		{name: "eighteen decimals", answer: "1000000000000000000", decimals: 18, wantPrice: "1.000000000000000000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			answer, ok := new(big.Int).SetString(tt.answer, 10)
			if !ok {
				t.Fatalf("invalid test answer %q", tt.answer)
			}
			if got := formatAnswer(answer, tt.decimals); got != tt.wantPrice {
				t.Fatalf("formatAnswer() = %q, want %q", got, tt.wantPrice)
			}
		})
	}
}
