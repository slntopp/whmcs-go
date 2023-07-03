package whmcs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c, err := NewClient("test", "test", "test", false)
	assert.NoError(t, err)
	assert.NotNil(t, c.Authentication)
	assert.NotNil(t, c.Support)
	assert.NotNil(t, c.Tickets)
	assert.NotNil(t, c.System)
}
