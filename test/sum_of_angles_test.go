package test

import (
	"codewar"
	"codewar/internal/assert"
	"testing"
)

func TestAngle(t *testing.T) {
	const method = "sum of angles"

	actual := codewar.Angle(3)
	if err := assert.Equal(method, 180, actual); err != nil {
		t.Errorf("wrong result; %s", err.Error())
	}

	actual = codewar.Angle(4)
	if err := assert.Equal(method, 360, actual); err != nil {
		t.Errorf("wrong result; %s", err.Error())
	}
}
