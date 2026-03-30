package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrimitivesParsing(t *testing.T) {
	var data = []byte{58, 240, 9, 68, 156, 196, 206, 97, 127, 255, 127, 255, 255, 255, 127, 255, 255, 255, 255, 255, 255, 255, 127, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 127, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 195, 245, 72, 64, 105, 87, 20, 139, 10, 191, 5, 64, 1}
	offset := 8

	i8, n := extractPrimitive(data, offset, "i8")
	assert.Equal(t, 127, i8)
	offset += n

	i16, n := extractPrimitive(data, offset, "i16")
	assert.Equal(t, 32767, i16)
	offset += n

	i32, n := extractPrimitive(data, offset, "i32")
	assert.Equal(t, 2147483647, i32)
	offset += n

	i64, n := extractPrimitive(data, offset, "i64")
	assert.Equal(t, 9223372036854775807, i64)
	offset += n

	i128, n := extractPrimitive(data, 23, "i128")
	assert.Equal(t, "170141183460469231731687303715884105727", i128)
	offset += n

	u8, n := extractPrimitive(data, offset, "u8")
	assert.Equal(t, 255, u8)
	offset += n

	u16, n := extractPrimitive(data, offset, "u16")
	assert.Equal(t, 65535, u16)
	offset += n

	u32, n := extractPrimitive(data, offset, "u32")
	assert.Equal(t, 4294967295, u32)
	offset += n

	u64, n := extractPrimitive(data, offset, "u64")
	assert.Equal(t, uint(18446744073709551615), u64)
	offset += n

	u128, n := extractPrimitive(data, offset, "u128")
	assert.Equal(t, "340282366920938463463374607431768211455", u128)
	offset += n

	f32, n := extractPrimitive(data, offset, "f32")
	assert.Equal(t, 3.14, f32)
	offset += n

	f64, n := extractPrimitive(data, offset, "f64")
	assert.Equal(t, 2.718281828459045, f64)
	offset += n

	bool, n := extractPrimitive(data, offset, "bool")
	assert.Equal(t, true, bool)
	offset += n
}

func TestBoxedValue(t *testing.T) {
	var data = []byte{159, 49, 222, 167, 196, 68, 160, 162, 30}
	offset := 8

	boxed, _ := extractPrimitive(data, offset, "u8")
	assert.Equal(t, 30, boxed)
}
