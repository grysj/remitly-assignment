package main

import (
	"fmt"

	parsejson "github.com/grysj/check/parse"
	readjson "github.com/grysj/check/read"
)

func Check(value interface{}) (bool, error) {
	var data map[string]interface{}
	var err error

	switch v := value.(type) {
	case string:
		data, err = readjson.ReadFile(v)
		if err != nil {
			return false, err
		}
	case []byte:
		data, err = readjson.ReadByte(v)
		if err != nil {
			return false, err
		}

	default:
		return false, fmt.Errorf("invalid type of value")
	}

	res, err := parsejson.GetResourseList(data)
	if err != nil {
		return false, err
	}

	_, ok := res["*"]

	return !ok, nil
}
