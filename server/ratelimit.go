package server

import (
	"crypto/sha256"
	"fmt"
	"net"
	"net/http"
	"net/netip"
	"strings"
	"time"

	"github.com/sethvargo/go-limiter"
	"github.com/sethvargo/go-limiter/httplimit"
	"github.com/uatu/config"
)

const (
	defaultRateLimitTokens   uint64 = 60
	defaultRateLimitInterval        = time.Minute
	rateLimitKeyPrefix              = "uatu:ratelimit:"
)

func newRateLimiter(cfg config.RateLimitConfig, store limiter.Store) (func(http.Handler) http.Handler, error) {
	if store == nil {
		return nil, fmt.Errorf("rate limit store is required")
	}

	if _, _, err := rateLimitOptions(cfg); err != nil {
		return nil, err
	}

	keyFunc, err := clientIPKeyFunc(cfg)
	if err != nil {
		return nil, err
	}

	middleware, err := httplimit.NewMiddleware(store, keyFunc)
	if err != nil {
		return nil, fmt.Errorf("create rate limit middleware: %w", err)
	}
	return middleware.Handle, nil
}

func rateLimitOptions(cfg config.RateLimitConfig) (uint64, time.Duration, error) {
	tokens := cfg.Tokens
	if tokens == 0 {
		tokens = defaultRateLimitTokens
	}
	interval := cfg.Interval
	if interval == 0 {
		interval = defaultRateLimitInterval
	}
	// go-redisstore stores its retention TTL in seconds. Reject shorter
	// intervals rather than silently creating counters that expire immediately.
	if interval < time.Second {
		return 0, 0, fmt.Errorf("rate limit interval must be at least one second")
	}
	return tokens, interval, nil
}

// clientIPKeyFunc resolves the client address then converts it to a Redis-safe
// rate-limit key. Proxy headers are considered only when the direct TCP peer
// is in the explicit allowlist. This prevents an internet client that reaches
// the service directly from minting buckets with a forged forwarding header.
func clientIPKeyFunc(cfg config.RateLimitConfig) (httplimit.KeyFunc, error) {
	header := http.CanonicalHeaderKey(strings.TrimSpace(cfg.TrustedHeader))
	prefixes, err := trustedProxyPrefixes(cfg.TrustedProxyCIDRs)
	if err != nil {
		return nil, err
	}
	if header == "" && len(prefixes) > 0 {
		return nil, fmt.Errorf("RATE_LIMIT_TRUSTED_PROXY_CIDRS requires RATE_LIMIT_TRUSTED_HEADER")
	}
	if header != "" && len(prefixes) == 0 {
		return nil, fmt.Errorf("RATE_LIMIT_TRUSTED_HEADER requires RATE_LIMIT_TRUSTED_PROXY_CIDRS")
	}

	return func(r *http.Request) (string, error) {
		remote, err := remoteAddrIP(r.RemoteAddr)
		if err != nil {
			return "", err
		}

		client := remote
		if header != "" && isTrustedProxy(remote, prefixes) {
			if header == "X-Forwarded-For" {
				client = clientIPFromXFF(r.Header.Values(header), remote, prefixes)
			} else {
				client = clientIPFromSingleHeader(r.Header.Values(header), remote)
			}
		}
		return rateLimitKey(client), nil
	}, nil
}

func rateLimitKey(ip netip.Addr) string {
	// IPv6 privacy addresses can rotate freely within a /64. Bucket that
	// network together before hashing so it cannot defeat a per-client limit.
	if ip.Is6() {
		ip = netip.PrefixFrom(ip, 64).Masked().Addr()
	}
	digest := sha256.Sum256([]byte(ip.String()))
	return fmt.Sprintf("%s%x", rateLimitKeyPrefix, digest)
}

func trustedProxyPrefixes(raw string) ([]netip.Prefix, error) {
	if strings.TrimSpace(raw) == "" {
		return nil, nil
	}

	parts := strings.Split(raw, ",")
	prefixes := make([]netip.Prefix, 0, len(parts))
	for _, part := range parts {
		prefix, err := netip.ParsePrefix(strings.TrimSpace(part))
		if err != nil {
			return nil, fmt.Errorf("invalid RATE_LIMIT_TRUSTED_PROXY_CIDRS entry %q: %w", part, err)
		}
		prefixes = append(prefixes, prefix.Masked())
	}
	return prefixes, nil
}

func remoteAddrIP(remoteAddr string) (netip.Addr, error) {
	host, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		host = remoteAddr
	}
	ip, err := netip.ParseAddr(host)
	if err != nil {
		return netip.Addr{}, fmt.Errorf("invalid remote address %q: %w", remoteAddr, err)
	}
	return ip.Unmap(), nil
}

func clientIPFromSingleHeader(values []string, fallback netip.Addr) netip.Addr {
	if len(values) == 0 {
		return fallback
	}
	// A proxy that overwrites a single-IP header is authoritative. If it
	// appends instead, the last value is from the nearest proxy. Invalid values
	// fail closed to the known direct peer rather than trusting an older value.
	ip, err := netip.ParseAddr(strings.TrimSpace(values[len(values)-1]))
	if err != nil {
		return fallback
	}
	return ip.Unmap()
}

func clientIPFromXFF(values []string, fallback netip.Addr, prefixes []netip.Prefix) netip.Addr {
	for headerIndex := len(values) - 1; headerIndex >= 0; headerIndex-- {
		value := values[headerIndex]
		for value != "" {
			entry := value
			if index := strings.LastIndexByte(value, ','); index >= 0 {
				entry, value = value[index+1:], value[:index]
			} else {
				value = ""
			}

			entry = strings.TrimSpace(entry)
			if entry == "" {
				continue
			}
			ip, err := netip.ParseAddr(entry)
			if err != nil {
				return fallback
			}
			ip = ip.Unmap()
			if !isTrustedProxy(ip, prefixes) {
				return ip
			}
		}
	}
	return fallback
}

func isTrustedProxy(ip netip.Addr, prefixes []netip.Prefix) bool {
	for _, prefix := range prefixes {
		if prefix.Contains(ip) {
			return true
		}
	}
	return false
}
