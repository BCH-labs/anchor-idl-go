package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrimitivesParsing(t *testing.T) {
	var data = []byte{58, 240, 9, 68, 156, 196, 206, 97, 246, 133, 255, 222, 207, 255, 255, 1, 0, 132, 226, 80, 108, 230, 252, 1, 0, 0, 0, 64, 34, 138, 9, 122, 196, 134, 90, 168, 76, 59, 203, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 195, 245, 72, 64, 105, 87, 20, 139, 10, 191, 5, 64, 1}
	offset := 8

	i8, n := extractPrimitive(data, offset, "i8")
	assert.Equal(t, int8(-10), i8.(int8))
	offset += n

	i16, n := extractPrimitive(data, offset, "i16")
	assert.Equal(t, int16(-123), i16.(int16))
	offset += n

	i32, n := extractPrimitive(data, offset, "i32")
	assert.Equal(t, int32(-12322), i32)
	offset += n

	i64, n := extractPrimitive(data, offset, "i64")
	assert.Equal(t, int64(-223372036854775807), i64)
	offset += n

	i128, n := extractPrimitive(data, offset, "i128")
	assert.Equal(t, "-70141183460469231731687303715884105727", i128)
	offset += n

	u8, n := extractPrimitive(data, offset, "u8")
	assert.Equal(t, uint8(255), u8)
	offset += n

	u16, n := extractPrimitive(data, offset, "u16")
	assert.Equal(t, uint16(65535), u16)
	offset += n

	u32, n := extractPrimitive(data, offset, "u32")
	assert.Equal(t, uint32(4294967295), u32)
	offset += n

	u64, n := extractPrimitive(data, offset, "u64")
	assert.Equal(t, uint64(18446744073709551615), u64)
	offset += n

	u128, n := extractPrimitive(data, offset, "u128")
	assert.Equal(t, "340282366920938463463374607431768211455", u128)
	offset += n

	f32, n := extractPrimitive(data, offset, "f32")
	assert.Equal(t, float32(3.14), f32)
	offset += n

	f64, n := extractPrimitive(data, offset, "f64")
	assert.Equal(t, float64(2.718281828459045), f64)
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
