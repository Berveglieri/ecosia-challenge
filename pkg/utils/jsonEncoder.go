package utils

import (
	"encoding/json"
	"net/http"
)

func ConvertToJsonWithKey(w http.ResponseWriter, status int, data interface{}, key string) error {
	wrapper := make(map[string]interface{})

	wrapper[key] = data

	jsonContent, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonContent)

	return nil

}

func ConvertToJson(w http.ResponseWriter, status int, data interface{}) error {
	jsonContent, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonContent)

	return nil

}
