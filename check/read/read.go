package read

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFile(filePath string) (map[string]interface{}, error) {

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

func ReadByte(byteValue []byte) (map[string]interface{}, error) {

	var data map[string]interface{}

	if !json.Valid(byteValue) {
		return data, fmt.Errorf("invalid json")
	}
	json.Unmarshal(byteValue, &data)

	return data, nil

}
