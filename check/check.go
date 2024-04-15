package checkJson

import (
	"fmt"
	parsejson "parseJson"
	readjson "readJson"
)

func Check(filepath string) (bool, error) {

	data, err := readjson.Read(filepath)
	if err != nil {
		return false, err
	}

	res, err := parsejson.GetResourseList(data)
	if err != nil {
		return false, err
	}

	for _, item := range res {
		switch toCheck := item.(type) {
		case string:
			if toCheck == "*" {
				return false, nil
			}
		default:
			continue
		}

	}

	return true, nil

}

func main() {

	fmt.Println("hello!")

}
