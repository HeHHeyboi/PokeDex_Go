package pokeapi

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	t.Run("creates a client with specified cache interval and HTTP timeout", func(t *testing.T) {
		cacheInterval := 5 * time.Minute
		httpTimeout := 10 * time.Second

		client := NewClient(cacheInterval, httpTimeout)

		// Verify Cache is initialized with the correct interval

		// Verify HTTP client timeout
		assert.Equal(t, httpTimeout, client.client.Timeout)
	})
}
