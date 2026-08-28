package feeds

import (
	"math/big"
	"strings"
	"testing"
)

func TestSupportedTokensSlugs(t *testing.T) {
	seen := make(map[string]int, countSlugs(SupportedTokens))

	for i, token := range SupportedTokens {
		if len(token.Slugs) == 0 {
			t.Errorf("SupportedTokens[%d] (%s) has no slugs", i, token.ChainSlug)
		}
		for _, slug := range token.Slugs {
			if slug != strings.ToLower(slug) {
				t.Errorf("slug %q is not lowercase, cache lookups are case sensitive", slug)
			}
			if first, ok := seen[slug]; ok {
				t.Errorf("slug %q is in SupportedTokens[%d] and [%d], its cache key would be overwritten", slug, first, i)
			}
			seen[slug] = i
		}
	}
}

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
