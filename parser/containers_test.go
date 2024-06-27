package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainerParsing(t *testing.T) {
	var data = []byte{114, 122, 122, 101, 12, 55, 37, 178, 13, 0, 0, 0, 72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33, 3, 0, 0, 0, 72, 105, 33, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 0}
	offset := 8

	st, n := extractPrimitive(data, offset, "string")
	assert.Equal(t, "Hello, World!", st)
	offset += n

	vec, n := extractVector(data, nil, offset, "bytes")
	assert.Equal(t, "Hi!", vec)
	offset += n

	arr, n := extractArray(data, nil, offset, []interface{}{"u8", 10})
	assert.Equal(t, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, arr)
	offset += n

	// opt, _ := extractOption(data, nil, offset, "u8")
	// assert.Equal(t, nil, opt)
}
