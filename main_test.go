package goretry

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	testCases := []struct {
		name    string
		timeout time.Duration
		wantErr bool
	}{
		{
			name:    "success",
			timeout: 10 * time.Second,
			wantErr: false,
		},
		{
			name:    "timeout exceeded",
			timeout: 1 * time.Millisecond,
			wantErr: true,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var attempt int
			err := Retry(tt.timeout, func() error {
				if attempt < 3 {
					attempt++
					fmt.Printf("attempt: %d \n", attempt)
					return errors.New("timeout exceeded")
				}
				return nil
			})
			if tt.wantErr {
				assert.EqualError(t, err, "timeout exceeded")
				t.Log(err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
