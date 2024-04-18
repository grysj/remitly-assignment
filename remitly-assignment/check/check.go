package check

import (
	"fmt"

	parse "github.com/grysj/remitly-assignment/remitly-assignment/parse"
	read "github.com/grysj/remitly-assignment/remitly-assignment/read"
)

func Check(value interface{}) (bool, error) {
	var data map[string]interface{}
	var err error

	switch v := value.(type) {
	case string:
		data, err = read.ReadFile(v)
		if err != nil {
			return false, err
		}
	case []byte:
		data, err = read.ReadByte(v)
		if err != nil {
			return false, err
		}

	default:
		return false, fmt.Errorf("invalid type of value")
	}

	res, err := parse.GetResourseList(data)
	if err != nil {
		return false, err
	}

	_, ok := res["*"]

	return !ok, nil
}
