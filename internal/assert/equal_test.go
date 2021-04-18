package assert

import (
	"testing"
)

func TestCompareIntegerType(t *testing.T) {
	// int
	if err := NotEqual("compare integer type", 1, 2); err != nil {
		t.Error(err)
	}
	if err := Equal("compare integer type", 1, 1); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare integer type", 1, 2); err != nil {
		t.Error(err)
	}
	if err := Equal("compare integer type", 1, 1); err != nil {
		t.Error(err)
	}
	// int8
	if err := NotEqual("compare integer type", int8(1), int8(2)); err != nil {
		t.Error(err)
	}
	if err := Equal("compare integer type", int8(1), int8(1)); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare integer type", 1, int8(1)); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare integer type", 0xFF>>1, int8(0xFF>>1)); err != nil {
		t.Error(err)
	}
	// int16
	if err := NotEqual("compare integer type", int16(1), int16(2)); err != nil {
		t.Error(err)
	}
	if err := Equal("compare integer type", int16(1), int16(1)); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare integer type", 1, int16(1)); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare integer type", 0xFFFF>>1, int16(0xFFFF>>1)); err != nil {
		t.Error(err)
	}
	// int32
	if err := NotEqual("compare integer type", int32(1), int32(2)); err != nil {
		t.Error(err)
	}
	if err := Equal("compare integer type", int32(1), int32(1)); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare integer type", 1, int32(1)); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare integer type", 0xFFFFFFFF>>1, int32(0xFFFFFFFF>>1)); err != nil {
		t.Error(err)
	}
	// int64
	if err := NotEqual("compare integer type", int64(1), int64(2)); err != nil {
		t.Error(err)
	}
	if err := Equal("compare integer type", int64(1), int64(1)); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare integer type", 1, int64(1)); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare integer type", 0xFFFFFFFFFFFFFFFF>>1, int64(0xFFFFFFFFFFFFFFFF>>1)); err != nil {
		t.Error(err)
	}
	// uint8
	if err := NotEqual("compare integer type", uint8(1), uint8(2)); err != nil {
		t.Error(err)
	}
	if err := Equal("compare integer type", uint8(1), uint8(1)); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare integer type", 1, uint(1)); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare integer type", 0xFF, uint8(0xFF)); err != nil {
		t.Error(err)
	}
	// uint16
	if err := NotEqual("compare integer type", uint16(1), uint16(2)); err != nil {
		t.Error(err)
	}
	if err := Equal("compare integer type", uint16(1), uint16(1)); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare integer type", 1, uint16(1)); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare integer type", 0xFFFF, uint16(0xFFFF)); err != nil {
		t.Error(err)
	}
	// uint32
	if err := NotEqual("compare integer type", uint32(1), uint32(2)); err != nil {
		t.Error(err)
	}
	if err := Equal("compare integer type", uint32(1), uint32(1)); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare integer type", 1, uint32(1)); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare integer type", 0xFFFFFFFF, uint32(0xFFFFFFFF)); err != nil {
		t.Error(err)
	}
	// uint64
	if err := NotEqual("compare integer type", uint64(1), uint64(2)); err != nil {
		t.Error(err)
	}
	if err := Equal("compare integer type", uint64(1), uint64(1)); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare integer type", 1, uint64(1)); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare integer type", 0xFFFFFFFFFFFFFFF, uint64(0xFFFFFFFFFFFFFFF)); err != nil {
		t.Error(err)
	}
}

func TestCompareFloatType(t *testing.T) {
	// float32
	if err := NotEqual("compare integer type", float32(1.0), float32(2.0)); err != nil {
		t.Error(err)
	}
	if err := Equal("compare integer type", float32(1.1), float32(1.1)); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare integer type", float32(1.2), float32(2.2)); err != nil {
		t.Error(err)
	}
	if err := Equal("compare integer type", float32(1.2), float32(1.2)); err != nil {
		t.Error(err)
	}
	// float64
	if err := NotEqual("compare integer type", float64(1.0), float64(2.0)); err != nil {
		t.Error(err)
	}
	if err := Equal("compare integer type", float64(1.1), float64(1.1)); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare integer type", float64(1.2), float64(2.2)); err != nil {
		t.Error(err)
	}
	if err := Equal("compare integer type", float64(1.2), float64(1.2)); err != nil {
		t.Error(err)
	}
}

func TestCompareStringType(t *testing.T) {
	if err := NotEqual("compare string type", "1", "2"); err != nil {
		t.Error(err)
	}
	if err := Equal("compare string type", "1", "1"); err != nil {
		t.Error(err)
	}
}

func TestCompareDifferentType(t *testing.T) {
	// string and integer
	if err := NotEqual("compare string and integer type", "1", 0); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare string and integer type", 1, "0"); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare string and integer type", 0, "0"); err != nil {
		t.Error(err)
	}
	// string and float32
	if err := NotEqual("compare string and float32 type", "1.0", float32(1.1)); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare string and float32 type", float32(1.2), "1.0"); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare string and float32 type", float32(0.0), "0.0"); err != nil {
		t.Error(err)
	}
	// string and float64
	if err := NotEqual("compare string and float64 type", "1.0", float64(1.1)); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare string and float64 type", float64(1.2), "1.0"); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare string and float64 type", float64(0.0), "0.0"); err != nil {
		t.Error(err)
	}
	// string and bool
	if err := NotEqual("compare string and bool type", "true", true); err != nil {
		t.Error(err)
	}
	if err := NotEqual("compare string and bool type", false, "false"); err != nil {
		t.Error(err)
	}
}
