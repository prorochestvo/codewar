package test

import (
	"codewar"
	"codewar/internal/assert"
	"testing"
)

func TestCountingDuplicates(t *testing.T) {
	testCountingDuplicates(t, codewar.CountingDuplicates, "counting duplicates (v1)")
	testCountingDuplicates(t, codewar.CountingDuplicatesV2, "counting duplicates (v2)")
	testCountingDuplicates(t, codewar.CountingDuplicatesV3, "counting duplicates (v3)")
}

func testCountingDuplicates(t *testing.T, f func(s string) int, method string) {
	actual := f("abcde")
	if err := assert.Equal(method, 0, actual); err != nil {
		t.Errorf("wrong result; %s", err.Error())
	}

	actual = f("abcdea")
	if err := assert.Equal(method, 1, actual); err != nil {
		t.Errorf("wrong result; %s", err.Error())
	}

	actual = f("abcdeaB11")
	if err := assert.Equal(method, 3, actual); err != nil {
		t.Errorf("wrong result; %s", err.Error())
	}

	actual = f("indivisibility")
	if err := assert.Equal(method, 1, actual); err != nil {
		t.Errorf("wrong result; %s", err.Error())
	}
}

func BenchmarkCountingDuplicatesC1V1(b *testing.B) {
	benchmarkCountingDuplicates(b, codewar.CountingDuplicates, "indivisibility")
}

func BenchmarkCountingDuplicatesC1V2(b *testing.B) {
	benchmarkCountingDuplicates(b, codewar.CountingDuplicatesV2, "indivisibility")
}

func BenchmarkCountingDuplicatesC1V3(b *testing.B) {
	benchmarkCountingDuplicates(b, codewar.CountingDuplicatesV3, "indivisibility")
}

func BenchmarkCountingDuplicatesC3V1(b *testing.B) {
	benchmarkCountingDuplicates(b, codewar.CountingDuplicates, "abcdeaB11")
}

func BenchmarkCountingDuplicatesC3V2(b *testing.B) {
	benchmarkCountingDuplicates(b, codewar.CountingDuplicatesV2, "abcdeaB11")
}

func BenchmarkCountingDuplicatesC3V3(b *testing.B) {
	benchmarkCountingDuplicates(b, codewar.CountingDuplicatesV3, "abcdeaB11")
}

func benchmarkCountingDuplicates(b *testing.B, f func(s string) int, txt string) {
	for i := 0; i < b.N; i++ {
		f(txt)
	}
}
