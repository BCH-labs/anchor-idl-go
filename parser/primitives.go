package parser

import (
	"encoding/binary"

	"github.com/btcsuite/btcutil/base58"
)

func extractPrimitive(data []byte, offset int, argType string) (interface{}, int) {
	switch argType {
	//TODO: add case for vec, object
	case "u64":
		if len(data[offset:]) < 8 {
			return nil, 8
		} else {
			return binary.LittleEndian.Uint64(data[offset : offset+8]), 8
		}
	case "u32":
		if len(data[offset:]) < 4 {
			return nil, 4
		} else {
			return binary.LittleEndian.Uint32(data[offset : offset+4]), 4
		}
	case "u8":
		if len(data[offset:]) < 1 {
			return nil, 1
		} else {
			return data[offset], 1
		}
	case "bool":
		if len(data[offset:]) < 1 {
			return nil, 1
		} else {
			return data[offset], 1
		}
	case "publicKey":
		if len(data[offset:]) < 32 {
			return nil, 32
		} else {
			return base58.Encode(data[offset : offset+32]), 32
		}
	case "string":
		strLen := binary.LittleEndian.Uint32(data[offset : offset+4])
		var n int = 4
		if len(data[offset+n:]) < int(strLen) {
			return nil, n
		} else {
			return string(data[offset+n : offset+n+int(strLen)]), n + int(strLen)
		}
	case "bytes":
		return extractVector(data, nil, offset, "u8")
	}
	return nil, 0
}
