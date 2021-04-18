package test

import (
	"codewar"
	"codewar/internal/assert"
	"math"
	"testing"
)

func TestFindOutlier(t *testing.T) {
	const method = "find the parity outlier"

	actual := codewar.FindOutlier([]int{2, 6, 8, -10, 3})
	if err := assert.Equal(method, 3, actual); err != nil {
		t.Errorf("wrong result; %s", err.Error())
	}

	actual = codewar.FindOutlier([]int{206847684, 1056521, 7, 17, 1901, 21104421, 7, 1, 35521, 1, 7781})
	if err := assert.Equal(method, 206847684, actual); err != nil {
		t.Errorf("wrong result; %s", err.Error())
	}

	actual = codewar.FindOutlier([]int{math.MaxInt32, 0, 1})
	if err := assert.Equal(method, 0, actual); err != nil {
		t.Errorf("wrong result; %s", err.Error())
	}
}
