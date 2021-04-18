package test

import (
	"codewar"
	"codewar/internal/assert"
	"testing"
)

func TestHighestScoringWord(t *testing.T) {
	const method = "highest scoring word"

	actual := codewar.HighestScoringWord("man i need a taxi up to ubud")
	if err := assert.Equal(method, "taxi", actual); err != nil {
		t.Errorf("wrong result; %s", err.Error())
	}

	actual = codewar.HighestScoringWord("what time are we climbing up the volcano")
	if err := assert.Equal(method, "volcano", actual); err != nil {
		t.Errorf("wrong result; %s", err.Error())
	}

	actual = codewar.HighestScoringWord("take me to semynak")
	if err := assert.Equal(method, "semynak", actual); err != nil {
		t.Errorf("wrong result; %s", err.Error())
	}
}
