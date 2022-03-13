package rate_limit

import (
	"time"
)

const (
	RATELIMITER_ClearEveryN = 5
)

func rfc3339now() string {
	return time.Now().Format(time.RFC3339)
}
