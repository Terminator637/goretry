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
			var count int
			err := Retry(tt.timeout, func() error {
				if count < 3 {
					count++
					fmt.Println(count)
					return errors.New("timeout exceeded")
				}
				return nil
			})
			if tt.wantErr {
				assert.Error(t, err)
				t.Log(err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
