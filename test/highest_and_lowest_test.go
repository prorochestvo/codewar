package test

import (
	"codewar"
	"codewar/internal/assert"
	"testing"
)

func TestHighAndLow(t *testing.T) {
	const method = "highest and lowest"

	actual := codewar.HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4")
	if err := assert.Equal(method, "42 -9", actual); err != nil {
		t.Errorf("wrong result; %s", err.Error())
	}
}
