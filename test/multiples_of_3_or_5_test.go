package test

import (
	"codewar"
	"codewar/internal/assert"
	"testing"
)

func TestMultiple3And5(t *testing.T) {
	const method = "multiples of 3 or 5"

	actual := codewar.Multiple3And5(10)
	if err := assert.Equal(method, 23, actual); err != nil {
		t.Errorf("wrong result; %s", err.Error())
	}
}
