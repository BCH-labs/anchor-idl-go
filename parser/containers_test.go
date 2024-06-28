package parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainerParsing(t *testing.T) {
	var data = []byte{114, 122, 122, 101, 12, 55, 37, 178, 13, 0, 0, 0, 72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33, 3, 0, 0, 0, 72, 105, 33, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 0}
	offset := 8

	st, n := extractPrimitive(data, offset, "string")
	assert.Equal(t, "Hello, World!", st)
	offset += n

	// here call to extractPrimitive
	// because this method handles all types that are represented as string
	vec, n := extractPrimitive(data, offset, "bytes")
	expected := strings.Join([]string{"72", "105", "33"}, ", ")
	assert.Equal(t, expected, vec)
	offset += n

	arr, n := extractArray(data, nil, offset, []interface{}{"u8", 10})
	expected = strings.Join([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, ", ")
	assert.Equal(t, expected, arr)
	offset += n

	arg := make(map[string]interface{})
	arg["option"] = "u8"
	opt, _ := extractValue(data, nil, offset, arg)
	// IDK if option is nil than it is 0 anyway, weird
	assert.Equal(t, uint8(0), opt)
}
