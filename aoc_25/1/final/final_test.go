package final

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFinal(t *testing.T) {
	val, err := Run("../input")
	assert.Nil(t, err)
	assert.Equal(t, 5923, val)
}
