package math

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestAdder(t *testing.T) {
	a, b := 1, 2
	added := Adder(a, b)
	
	assert.Equal(t, 3, added)
}