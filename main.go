package goretry

import (
	"github.com/cenkalti/backoff"
	"time"
)

// Retry is a retry helper with timeout
func Retry(timeout time.Duration, op func() error) error {
	bo := backoff.NewExponentialBackOff()
	bo.MaxInterval = time.Second * 5
	bo.MaxElapsedTime = timeout
	return backoff.Retry(op, bo)
}
