package anchoridl

import (
	"encoding/json"
	"github.com/BCH-labs/anchor-idl-go/parser"
	"log"
)

func ParseData(data []byte, idlJson string) (map[string]interface{}, error) {
	var idl map[string]interface{}
	err := json.Unmarshal([]byte(idlJson), &idl)
	if err != nil {
		log.Fatal(err)
	}

	inst, err := parser.GetArgs(data, idl)
	if err != nil {
		log.Fatal(err)
	}
	return inst, nil
}
