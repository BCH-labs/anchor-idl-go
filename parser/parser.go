package parser

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/BCH-labs/anchor-idl-go/utils"
)

func GetArgs(data []byte, idl map[string]interface{}) (map[string]interface{}, error) {
	if idl == nil {
		return nil, errors.New("idl is nil")
	}
	if data == nil {
		return nil, errors.New("data is nil")
	}

	instructions, ok := idl["instructions"].([]interface{})
	if !ok {
		return nil, errors.New("instructions not found in IDL")
	}

	types, ok := idl["types"].([]interface{})
	if !ok {
		return nil, errors.New("types not found in IDL")
	}

	for _, instruction := range instructions {
		instructionMap, ok := instruction.(map[string]interface{})
		if !ok {
			continue
		}
		if discriminator, ok := instructionMap["discriminator"].([]interface{}); ok {
			if len(discriminator) == 4 {
				discriminatorBytes := make([]byte, 4)
				for i, val := range discriminator {
					discriminatorBytes[i] = byte(val.(float64))
				}

				if bytes.Equal(data[:4], discriminatorBytes) {
					return extractArgs(data[4:], instructionMap["args"].([]interface{}), types), nil
				}
			}
		} else {
			instructionName, ok := instructionMap["name"].(string)

			if !ok {
				continue
			}
			instructionName = utils.ToSnakeCase(instructionName)
			hash := sha256.Sum256([]byte(fmt.Sprintf("global:%s", instructionName)))

			if bytes.Equal(data[:8], hash[:8]) {
				return extractArgs(data[8:], instructionMap["args"].([]interface{}), types), nil
			}
		}
	}

	return nil, errors.New("can't find instruction")
}
func extractArgs(data []byte, args []interface{}, types []interface{}) map[string]interface{} {
	argsValues := make(map[string]interface{})
	offset := 0
	for _, arg := range args {
		argMap := arg.(map[string]interface{})
		argName := argMap["name"].(string)
		argType := argMap["type"]

		var n int
		argsValues[argName], n = extractValue(data, types, offset, argType)
		offset += n
	}
	return argsValues
}

func extractValue(data []byte, types []interface{}, offset int, argType interface{}) (interface{}, int) {
	pType, ok := argType.(string)
	if ok {
		return extractPrimitive(data, offset, pType)
	}

	npType, ok := argType.(map[string]interface{})
	if ok {
		return extractNonPrimitive(data, types, offset, npType)
	}

	return nil, 0
}

func extractNonPrimitive(data []byte, types []interface{}, offset int, argType map[string]interface{}) (interface{}, int) {
	vec, ok := argType["vec"]
	if ok {
		return extractVector(data, types, offset, vec)
	}
	arr, ok := argType["array"]
	if ok {
		return extractArray(data, types, offset, arr)
	}
	obj, ok := argType["defined"].(string)
	if ok {
		return extractObject(data, types, offset, obj)
	}
	return nil, 0
}

func extractObject(data []byte, types []interface{}, offset int, typeName string) (string, int) {
	fields, err := extractNeccesaryFields(types, typeName)
	if err != nil {
		return "", 0
	}

	res := make(map[string]interface{})
	var n int = 0

	var n_i int
	for _, field := range fields {
		field, ok := field.(map[string]interface{})
		if !ok {
			log.Println("cannot cast field to map[string]interface{}, in extractObject")
		}
		res[field["name"].(string)], n_i = extractValue(data, types, offset+n, field["type"])
		n += n_i
	}

	json, err := json.Marshal(res)

	return string(json), n
}

func extractNeccesaryFields(types []interface{}, typeName string) ([]interface{}, error) {
	for _, t := range types {
		t, ok := t.(map[string]interface{})
		if !ok {
			return nil, errors.New("cannot cast type to map[string]interface{}")
		}
		if strings.EqualFold(t["name"].(string), typeName) {
			return t["type"].(map[string]interface{})["fields"].([]interface{}), nil
		}
	}
	return nil, errors.New(fmt.Sprintf("couldn't find type: %s, in: %+v", typeName, types))
}
