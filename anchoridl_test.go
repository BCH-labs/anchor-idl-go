package anchoridl_test

import (
	"strings"
	"testing"

	aidl "github.com/BCH-labs/anchor-idl-go"
)

var dataSample = []byte{158, 246, 177, 28, 124, 97, 55, 174, 2, 0, 0, 0, 5, 0, 0, 0, 72, 111, 109, 101, 114, 24, 9, 0, 0, 0, 72, 111, 109, 101, 114, 105, 116, 116, 97, 20, 3, 0, 0, 0, 72, 105, 33, 3, 0, 0, 0, 1, 0, 0, 0, 2, 0, 0, 0, 3, 0, 0, 0}
var programPubKey string = "4Sw3JdbtpW7xN9Rw4UbVY19o2v1LyoCvebnKPFu5fDr3"

var idlSample string = `{
  "version": "0.1.0",
  "name": "custom_struct_as_arg",
  "instructions": [
    {
      "name": "initialize",
      "accounts": [],
      "args": []
    },
    {
      "name": "obj",
      "accounts": [],
      "args": [
        {
          "name": "info",
          "type": {
            "defined": "Info"
          }
        }
      ]
    },
    {
      "name": "primitives",
      "accounts": [],
      "args": [
        {
          "name": "age",
          "type": "u8"
        },
        {
          "name": "isHuman",
          "type": "bool"
        },
        {
          "name": "phoneNumber",
          "type": "string"
        }
      ]
    },
    {
      "name": "vecs",
      "accounts": [],
      "args": [
        {
          "name": "familyMembers",
          "type": {
            "vec": {
              "defined": "FamilyData"
            }
          }
        },
        {
          "name": "favoriteBytes",
          "type": "bytes"
        },
        {
          "name": "favoriteNumbers",
          "type": {
            "vec": "u32"
          }
        }
      ]
    },
    {
      "name": "arrs",
      "accounts": [],
      "args": [
        {
          "name": "githubAccounts",
          "type": {
            "array": [
              {
                "defined": "GitHubInfo"
              },
              2
            ]
          }
        }
      ]
    }
  ],
  "types": [
    {
      "name": "Info",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "name",
            "type": "string"
          },
          {
            "name": "surname",
            "type": "string"
          },
          {
            "name": "location",
            "type": {
              "defined": "LocationInfo"
            }
          }
        ]
      }
    },
    {
      "name": "LocationInfo",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "city",
            "type": "string"
          },
          {
            "name": "postalCode",
            "type": {
              "array": [
                "u8",
                3
              ]
            }
          }
        ]
      }
    },
    {
      "name": "FamilyData",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "name",
            "type": "string"
          },
          {
            "name": "age",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "GitHubInfo",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "username",
            "type": "string"
          },
          {
            "name": "repos",
            "type": "u8"
          }
        ]
      }
    }
  ],
  "metadata": {
    "address": "4Sw3JdbtpW7xN9Rw4UbVY19o2v1LyoCvebnKPFu5fDr3"
  }
}
`

func TestVectorParsing(t *testing.T) {
	inst, err := aidl.ParseData(dataSample, idlSample)
	if err != nil {
		t.Fatal(err)
	}

	fMembers, ok := inst["familyMembers"].(string)
	if !ok {
		t.Fatal("familyMembers is not a string")
	}
	if !strings.EqualFold(fMembers, `{"age":24,"name":"Homer"}, {"age":20,"name":"Homeritta"}`) {
		t.Fatalf(`expected:'{"age":24,"name":"Homer"}, {"age":20,"name":"Homeritta"}', got: %s`,
			fMembers)
	}
	fb, ok := inst["favoriteBytes"].(string)
	if !ok {
		t.Fatal("favoriteBytes is not a string")
	}

	if !strings.EqualFold(fb, `72, 105, 33`) {
		t.Fatalf("expected:'72, 105, 33', got:%s", fb)
	}

	fn, ok := inst["favoriteNumbers"].(string)
	if !ok {
		t.Fatal("favoriteNumbers is not a string")
	}
	if !strings.EqualFold(fn, `1, 2, 3`) {
		t.Fatalf("expected: '1, 2, 3', got:%s", fn)
	}
}
