package modHello

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHi(t *testing.T) {
	assert.Equal(t, "Hello world", Hi(""))

	cases := []string{"a", "test", "Mod015", "Hello", "123"}
	for _, expect := range cases {
		assert.Equal(t, fmt.Sprintf("Hello %s", expect), Hi(expect))
	}
}
