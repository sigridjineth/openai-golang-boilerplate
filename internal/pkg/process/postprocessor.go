package process

import (
	"encoding/json"
	"fmt"
	"strings"
)

func TwoDimensionalInterfacesToJSONString(interfaces [][]interface{}) (string, error) {
	// Create a slice of maps that corresponds to the two-dimensional slice of interface{} values
	var data []map[string]interface{}
	for _, row := range interfaces {
		item := make(map[string]interface{})
		for i, value := range row {
			key := fmt.Sprintf("results%d", i+1)
			item[key] = value
		}
		data = append(data, item)
	}

	// Encode the slice of maps as a JSON-encoded byte slice
	resultJSON, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(resultJSON), nil
}

func TwoDimensionalInterfacesToJSONBytes(interfaces [][]interface{}) ([]byte, error) {
	// Create a slice of maps that corresponds to the two-dimensional slice of interface{} values
	var data []map[string]interface{}
	for _, row := range interfaces {
		item := make(map[string]interface{})
		for i, value := range row {
			key := fmt.Sprintf("results%d", i+1)
			item[key] = value
		}
		data = append(data, item)
	}

	// Encode the slice of maps as a JSON-encoded byte slice
	resultJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return resultJSON, nil
}

func ConvertNewLineToSpace(input string) string {
	return strings.ReplaceAll(input, "\n", " ")
}
