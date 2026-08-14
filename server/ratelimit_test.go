package server

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/sethvargo/go-limiter"
	"github.com/sethvargo/go-limiter/memorystore"
	"github.com/uatu/config"
)

type failingStore struct{}

func (failingStore) Take(context.Context, string) (uint64, uint64, uint64, bool, error) {
	return 0, 0, 0, false, errors.New("Redis credential leaked")
}

func (failingStore) Get(context.Context, string) (uint64, uint64, error) {
	return 0, 0, errors.New("unexpected get")
}

func (failingStore) Set(context.Context, string, uint64, time.Duration) error {
	return errors.New("unexpected set")
}

func (failingStore) Burst(context.Context, string, uint64) error {
	return errors.New("unexpected burst")
}

func (failingStore) Close(context.Context) error {
	return nil
}

func okHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

func request(t *testing.T, h http.Handler, remoteAddr string, headers map[string]string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodGet, "/quotes", nil)
	req.RemoteAddr = remoteAddr
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

func testRateLimiter(t *testing.T, cfg config.RateLimitConfig) func(http.Handler) http.Handler {
	t.Helper()
	tokens, interval, err := rateLimitOptions(cfg)
	if err != nil {
		t.Fatalf("rateLimitOptions() error = %v", err)
	}
	store, err := memorystore.New(&memorystore.Config{Tokens: tokens, Interval: interval})
	if err != nil {
		t.Fatalf("memorystore.New() error = %v", err)
	}
	t.Cleanup(func() { _ = store.Close(context.Background()) })

	limit, err := newRateLimiter(cfg, store)
	if err != nil {
		t.Fatalf("newRateLimiter() error = %v", err)
	}
	return limit
}

func testStore(t *testing.T) limiter.Store {
	t.Helper()
	store, err := memorystore.New(nil)
	if err != nil {
		t.Fatalf("memorystore.New() error = %v", err)
	}
	t.Cleanup(func() { _ = store.Close(context.Background()) })
	return store
}

func TestRateLimiterAllowsUpToBudgetThenRejects(t *testing.T) {
	limit := testRateLimiter(t, config.RateLimitConfig{Tokens: 3, Interval: time.Minute})
	handler := limit(okHandler())

	for i := 1; i <= 3; i++ {
		rec := request(t, handler, "10.0.0.1:5000", nil)
		if rec.Code != http.StatusOK {
			t.Fatalf("request %d: code = %d, want 200", i, rec.Code)
		}
		if got := rec.Header().Get("X-RateLimit-Limit"); got != "3" {
			t.Fatalf("request %d: X-RateLimit-Limit = %q, want 3", i, got)
		}
	}

	rec := request(t, handler, "10.0.0.1:5000", nil)
	if rec.Code != http.StatusTooManyRequests {
		t.Fatalf("fourth request: code = %d, want 429", rec.Code)
	}
	if got := rec.Header().Get("Retry-After"); got == "" {
		t.Fatal("fourth request: Retry-After is empty")
	}
	if got := rec.Header().Get("X-RateLimit-Remaining"); got != "0" {
		t.Fatalf("fourth request: X-RateLimit-Remaining = %q, want 0", got)
	}
}

func TestRateLimiterBudgetsPerClient(t *testing.T) {
	limit := testRateLimiter(t, config.RateLimitConfig{Tokens: 1, Interval: time.Minute})
	handler := limit(okHandler())

	if rec := request(t, handler, "10.0.0.1:5000", nil); rec.Code != http.StatusOK {
		t.Fatalf("first client: code = %d, want 200", rec.Code)
	}
	if rec := request(t, handler, "10.0.0.1:5000", nil); rec.Code != http.StatusTooManyRequests {
		t.Fatalf("first client repeat: code = %d, want 429", rec.Code)
	}
	// A different address must not inherit the exhausted budget.
	if rec := request(t, handler, "10.0.0.2:5000", nil); rec.Code != http.StatusOK {
		t.Fatalf("second client: code = %d, want 200", rec.Code)
	}
}

func TestRateLimiterIgnoresForwardedHeaderUnlessTrusted(t *testing.T) {
	limit := testRateLimiter(t, config.RateLimitConfig{Tokens: 1, Interval: time.Minute})
	handler := limit(okHandler())

	if rec := request(t, handler, "10.0.0.1:5000", nil); rec.Code != http.StatusOK {
		t.Fatalf("first request: code = %d, want 200", rec.Code)
	}
	// Same source address, forged header: the budget must still be spent.
	rec := request(t, handler, "10.0.0.1:5000", map[string]string{"X-Forwarded-For": "1.2.3.4"})
	if rec.Code != http.StatusTooManyRequests {
		t.Fatalf("forged header bypassed the limit: code = %d, want 429", rec.Code)
	}
}

func TestRateLimiterUsesTrustedHeaderWhenConfigured(t *testing.T) {
	limit := testRateLimiter(t, config.RateLimitConfig{
		Tokens:            1,
		Interval:          time.Minute,
		TrustedHeader:     "X-Forwarded-For",
		TrustedProxyCIDRs: "10.0.0.0/8",
	})
	handler := limit(okHandler())

	// One proxied address exhausts its own budget...
	if rec := request(t, handler, "10.0.0.1:5000", map[string]string{"X-Forwarded-For": "1.2.3.4"}); rec.Code != http.StatusOK {
		t.Fatalf("first proxied request: code = %d, want 200", rec.Code)
	}
	if rec := request(t, handler, "10.0.0.1:5000", map[string]string{"X-Forwarded-For": "1.2.3.4"}); rec.Code != http.StatusTooManyRequests {
		t.Fatalf("second proxied request: code = %d, want 429", rec.Code)
	}
	// ...while a different client behind the same proxy is unaffected.
	if rec := request(t, handler, "10.0.0.1:5000", map[string]string{"X-Forwarded-For": "5.6.7.8"}); rec.Code != http.StatusOK {
		t.Fatalf("other proxied client: code = %d, want 200", rec.Code)
	}
}

// Railway and similar platforms append the caller's address to X-Forwarded-For,
// so the entry the proxy wrote is the last one. Anything the client put in front
// of it must not change the bucket.
func TestRateLimiterKeysOnRightmostForwardedAddress(t *testing.T) {
	limit := testRateLimiter(t, config.RateLimitConfig{
		Tokens:            1,
		Interval:          time.Minute,
		TrustedHeader:     "X-Forwarded-For",
		TrustedProxyCIDRs: "10.0.0.0/8",
	})
	handler := limit(okHandler())

	if rec := request(t, handler, "10.0.0.9:5000", map[string]string{
		"X-Forwarded-For": "1.2.3.4",
	}); rec.Code != http.StatusOK {
		t.Fatalf("first request: code = %d, want 200", rec.Code)
	}
	// Same real client, but prefixing a forged hop must not mint a new budget.
	if rec := request(t, handler, "10.0.0.9:5000", map[string]string{
		"X-Forwarded-For": "9.9.9.9, 1.2.3.4",
	}); rec.Code != http.StatusTooManyRequests {
		t.Fatalf("forged prefix bypassed the limit: code = %d, want 429", rec.Code)
	}
}

// A client can send its own header line; the proxy's is appended after it.
func TestRateLimiterIgnoresForgedDuplicateHeader(t *testing.T) {
	limit := testRateLimiter(t, config.RateLimitConfig{
		Tokens:            1,
		Interval:          time.Minute,
		TrustedHeader:     "X-Forwarded-For",
		TrustedProxyCIDRs: "10.0.0.0/8",
	})
	handler := limit(okHandler())

	send := func() int {
		req := httptest.NewRequest(http.MethodGet, "/quotes", nil)
		req.RemoteAddr = "10.0.0.9:5000"
		req.Header.Add("X-Forwarded-For", "9.9.9.9") // forged, arrives first
		req.Header.Add("X-Forwarded-For", "1.2.3.4") // written by the proxy
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
		return rec.Code
	}

	if code := send(); code != http.StatusOK {
		t.Fatalf("first request: code = %d, want 200", code)
	}
	if code := send(); code != http.StatusTooManyRequests {
		t.Fatalf("second request: code = %d, want 429", code)
	}
}

func TestRateLimiterFallsBackToRemoteAddrWhenHeaderAbsent(t *testing.T) {
	limit := testRateLimiter(t, config.RateLimitConfig{
		Tokens:            1,
		Interval:          time.Minute,
		TrustedHeader:     "X-Forwarded-For",
		TrustedProxyCIDRs: "10.0.0.0/8",
	})
	handler := limit(okHandler())

	if rec := request(t, handler, "10.0.0.1:5000", nil); rec.Code != http.StatusOK {
		t.Fatalf("first request: code = %d, want 200", rec.Code)
	}
	if rec := request(t, handler, "10.0.0.1:5000", nil); rec.Code != http.StatusTooManyRequests {
		t.Fatalf("second request: code = %d, want 429", rec.Code)
	}
	if rec := request(t, handler, "10.0.0.2:5000", nil); rec.Code != http.StatusOK {
		t.Fatalf("other address: code = %d, want 200", rec.Code)
	}
}

// Single-value headers carry no chain, so the whole value is the key.
func TestRateLimiterHandlesSingleValueHeader(t *testing.T) {
	limit := testRateLimiter(t, config.RateLimitConfig{
		Tokens:            1,
		Interval:          time.Minute,
		TrustedHeader:     "CF-Connecting-IP",
		TrustedProxyCIDRs: "10.0.0.0/8",
	})
	handler := limit(okHandler())

	if rec := request(t, handler, "10.0.0.9:5000", map[string]string{
		"CF-Connecting-IP": "1.2.3.4",
	}); rec.Code != http.StatusOK {
		t.Fatalf("first request: code = %d, want 200", rec.Code)
	}
	if rec := request(t, handler, "10.0.0.9:5000", map[string]string{
		"CF-Connecting-IP": "1.2.3.4",
	}); rec.Code != http.StatusTooManyRequests {
		t.Fatalf("second request: code = %d, want 429", rec.Code)
	}
	if rec := request(t, handler, "10.0.0.9:5000", map[string]string{
		"CF-Connecting-IP": "5.6.7.8",
	}); rec.Code != http.StatusOK {
		t.Fatalf("second client: code = %d, want 200", rec.Code)
	}
}

func TestRateLimiterDefaultsWhenUnset(t *testing.T) {
	limit := testRateLimiter(t, config.RateLimitConfig{})
	rec := request(t, limit(okHandler()), "10.0.0.1:5000", nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("code = %d, want 200", rec.Code)
	}
	if got := rec.Header().Get("X-RateLimit-Limit"); got != "60" {
		t.Fatalf("X-RateLimit-Limit = %q, want the 60 default", got)
	}
}

func TestRateLimiterRequiresProxyAllowlistForTrustedHeader(t *testing.T) {
	_, err := newRateLimiter(config.RateLimitConfig{
		TrustedHeader: "X-Forwarded-For",
	}, testStore(t))
	if err == nil || !strings.Contains(err.Error(), "RATE_LIMIT_TRUSTED_PROXY_CIDRS") {
		t.Fatalf("newRateLimiter() error = %v, want missing proxy allowlist error", err)
	}
}

func TestRateLimiterRejectsProxyAllowlistWithoutHeader(t *testing.T) {
	_, err := newRateLimiter(config.RateLimitConfig{
		TrustedProxyCIDRs: "10.0.0.0/8",
	}, testStore(t))
	if err == nil || !strings.Contains(err.Error(), "RATE_LIMIT_TRUSTED_HEADER") {
		t.Fatalf("newRateLimiter() error = %v, want missing header error", err)
	}
}

func TestRateLimiterRejectsNegativeInterval(t *testing.T) {
	_, err := newRateLimiter(config.RateLimitConfig{
		Interval: -time.Second,
	}, testStore(t))
	if err == nil || !strings.Contains(err.Error(), "interval must be at least one second") {
		t.Fatalf("newRateLimiter() error = %v, want invalid interval error", err)
	}
}

func TestRateLimiterFailsClosedWhenCounterFails(t *testing.T) {
	limit, err := newRateLimiter(config.RateLimitConfig{}, failingStore{})
	if err != nil {
		t.Fatalf("newRateLimiter() error = %v", err)
	}

	rec := request(t, limit(okHandler()), "10.0.0.1:5000", nil)
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("code = %d, want 500", rec.Code)
	}
	if strings.Contains(rec.Body.String(), "credential") {
		t.Fatalf("response leaked counter error: %q", rec.Body.String())
	}
}

func TestRateLimiterIgnoresHeadersFromUntrustedPeer(t *testing.T) {
	limit := testRateLimiter(t, config.RateLimitConfig{
		Tokens:            1,
		Interval:          time.Minute,
		TrustedHeader:     "X-Forwarded-For",
		TrustedProxyCIDRs: "10.0.0.0/8",
	})
	handler := limit(okHandler())

	if rec := request(t, handler, "192.0.2.9:5000", map[string]string{
		"X-Forwarded-For": "1.2.3.4",
	}); rec.Code != http.StatusOK {
		t.Fatalf("first request: code = %d, want 200", rec.Code)
	}
	if rec := request(t, handler, "192.0.2.9:5000", map[string]string{
		"X-Forwarded-For": "5.6.7.8",
	}); rec.Code != http.StatusTooManyRequests {
		t.Fatalf("untrusted peer header bypassed the limit: code = %d, want 429", rec.Code)
	}
}

func TestRateLimiterCanonicalizesIPv6To64(t *testing.T) {
	limit := testRateLimiter(t, config.RateLimitConfig{Tokens: 1, Interval: time.Minute})
	handler := limit(okHandler())

	if rec := request(t, handler, "[2001:db8:1234:5678::1]:5000", nil); rec.Code != http.StatusOK {
		t.Fatalf("first address: code = %d, want 200", rec.Code)
	}
	if rec := request(t, handler, "[2001:db8:1234:5678::2]:5000", nil); rec.Code != http.StatusTooManyRequests {
		t.Fatalf("second address in /64 bypassed the limit: code = %d, want 429", rec.Code)
	}
	if rec := request(t, handler, "[2001:db8:1234:5679::1]:5000", nil); rec.Code != http.StatusOK {
		t.Fatalf("address in a different /64: code = %d, want 200", rec.Code)
	}
}

func TestRateLimiterFailsClosedForMalformedForwardingHeader(t *testing.T) {
	limit := testRateLimiter(t, config.RateLimitConfig{
		Tokens:            1,
		Interval:          time.Minute,
		TrustedHeader:     "X-Forwarded-For",
		TrustedProxyCIDRs: "10.0.0.0/8",
	})
	handler := limit(okHandler())

	if rec := request(t, handler, "10.0.0.9:5000", map[string]string{
		"X-Forwarded-For": "not-an-ip",
	}); rec.Code != http.StatusOK {
		t.Fatalf("first malformed request: code = %d, want 200", rec.Code)
	}
	if rec := request(t, handler, "10.0.0.9:5000", map[string]string{
		"X-Forwarded-For": "still-not-an-ip",
	}); rec.Code != http.StatusTooManyRequests {
		t.Fatalf("malformed headers minted a fresh budget: code = %d, want 429", rec.Code)
	}
}
