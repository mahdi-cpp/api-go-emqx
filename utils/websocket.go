package utils

import (
	"encoding/json"
	"github.com/mahdi-cpp/api-go-emqx/model"
)

// ConvertObjectToBytes converts an Object to a JSON byte slice
func ConvertObjectToBytes(obj model.Object) ([]byte, error) {
	// Marshal the object to JSON
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return nil, err // Return nil and the error if marshaling fails
	}
	return jsonData, nil // Return the JSON byte slice
}
