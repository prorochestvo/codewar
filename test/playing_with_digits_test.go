package test

import (
	"codewar"
	"codewar/internal/assert"
	"testing"
)

func TestDigPow(t *testing.T) {
	const method = "playing with digits"

	actual := codewar.DigPow(89, 1)
	if err := assert.Equal(method, 1, actual); err != nil {
		t.Errorf("wrong result; %s", err.Error())
	}

	actual = codewar.DigPow(92, 1)
	if err := assert.Equal(method, -1, actual); err != nil {
		t.Errorf("wrong result; %s", err.Error())
	}
}
