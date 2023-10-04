package steg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHide(t *testing.T) {
	err := Hide(&HideConfig{}, OutputDebug)
	assert.NoError(t, err)
}
