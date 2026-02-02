package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const defaultWriteRateLimitPerMinute = 60

type writeRateLimiter struct {
	mu       sync.Mutex
	limit    int
	window   time.Duration
	counters map[string]*rateLimitCounter
}

type rateLimitCounter struct {
	count   int
	resetAt time.Time
}

func NewWriteRateLimiter() *writeRateLimiter {
	return &writeRateLimiter{
		limit:    readWriteRateLimit(),
		window:   time.Minute,
		counters: make(map[string]*rateLimitCounter),
	}
}

func (l *writeRateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := strings.TrimSpace(r.Header.Get("X-API-Key"))
		if apiKey == "" {
			next.ServeHTTP(w, r)
			return
		}
		expectedKey, err := writeAPIKey()
		if err != nil || apiKey != expectedKey {
			next.ServeHTTP(w, r)
			return
		}
		if !l.allow(apiKey) {
			writeJSONError(
				w,
				http.StatusTooManyRequests,
				"rate_limited",
				fmt.Sprintf("write rate limit exceeded (%d/min); please try again later", l.limit),
			)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (l *writeRateLimiter) allow(key string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	counter, ok := l.counters[key]
	if !ok || now.After(counter.resetAt) {
		counter = &rateLimitCounter{resetAt: now.Add(l.window)}
		l.counters[key] = counter
	}
	if counter.count >= l.limit {
		return false
	}
	counter.count++
	return true
}

func readWriteRateLimit() int {
	if raw := strings.TrimSpace(os.Getenv("RATE_LIMIT_WRITE_PER_MINUTE")); raw != "" {
		value, err := strconv.Atoi(raw)
		if err == nil && value > 0 {
			return value
		}
	}
	return defaultWriteRateLimitPerMinute
}
