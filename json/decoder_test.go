package learn_golang_json

import (
	"os"
	"encoding/json"
	"testing"
	"fmt"
)

func TestDecoder(t *testing.T) {
	reader, _ := os.Open("sample.json")
	decoder 	:= json.NewDecoder(reader)

	// fmt.Println(reader)
	// fmt.Println(decoder)

	var data map[string]interface{}
	decoder.Decode(&data)
	
	fmt.Println(data)
}

