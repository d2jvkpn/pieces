package rate_limit

import (
	"context"
	"time"
)

const (
	RATELIMITER_ClearEveryN = 5
)

func rfc3339now() string {
	return time.Now().Format(time.RFC3339)
}

type Limiter interface {
	New(time.Duration, int) (Limiter, error)
	Allow(time.Time) bool
	AllowWithContext(context.Context, time.Time) bool
	Stop()
	Last() time.Time
}
