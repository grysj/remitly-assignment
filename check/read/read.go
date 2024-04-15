package readJson

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read(filePath string) (map[string]interface{}, error) {

	byteValue, err := os.ReadFile(filePath)
	var data map[string]interface{}
	if err != nil {
		return data, err
	}

	if !json.Valid(byteValue) {
		return data, fmt.Errorf("invalid json")
	}
	json.Unmarshal(byteValue, &data)

	return data, nil

}
