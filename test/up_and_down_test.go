package test

import (
	"codewar"
	"codewar/internal/assert"
	"testing"
)

func TestUpAndDown(t *testing.T) {
	const method = "highest scoring word"

	actual := codewar.UpAndDown("who hit retaining The That a we taken")
	if err := assert.Equal(method, "who RETAINING hit THAT a THE we TAKEN", actual); err != nil {
		t.Errorf("wrong result; %s", err.Error())
	}

	actual = codewar.UpAndDown("on I came up were so grandmothers")
	if err := assert.Equal(method, "i CAME on WERE up GRANDMOTHERS so", actual); err != nil {
		t.Errorf("wrong result; %s", err.Error())
	}

	actual = codewar.UpAndDown("way the my wall them him")
	if err := assert.Equal(method, "way THE my WALL him THEM", actual); err != nil {
		t.Errorf("wrong result; %s", err.Error())
	}
}
