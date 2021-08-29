package learn_golang_json

import (
	"os"
	"encoding/json"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("sample_encode.json")
	encoder 	:= json.NewEncoder(writer)

	product := map[string]interface{}{
		"id": 				"FE-0001",
		"name": 			"Flat Enterprise",
		"price":			250000,
		"image_url":	"https://google.com",
	}

	err := encoder.Encode(product)
	if err != nil {
		panic(err)
	}
}

