package learn_golang_json

import (
	"encoding/json"
	"testing"
	"fmt"
)

func TestJsonDecode(t *testing.T) {
	jsonString 	:= `{"FirstName":"John","LastName":"Doe","Age":20,"IsStudent":false,"Hobbies":["coding","reading"],"Addresses":[{"Street":"Street avenue","Country":"Swiss","PostalCode":"72111"},{"Street":"Street avenue 2","Country":"Berlin","PostalCode":"90213"}]}`
	jsonBytes 	:= []byte(jsonString)

	// type Person in file encode_json_test.go
	person := &Person{}

	err := json.Unmarshal(jsonBytes, person)

	if err != nil {
		panic(err)
	}

	fmt.Println(person)
	fmt.Println(person.FirstName)
	fmt.Println(person.LastName)
	fmt.Println(person.Age)
	fmt.Println(person.IsStudent)
	fmt.Println(person.Hobbies)
	fmt.Println(person.Addresses)

	fmt.Println("============================== JSON ARRAY ==========================")
	// JSON ARRAY
	
	jsonAddress := `[{"Street":"Street avenue","Country":"Swiss","PostalCode":"72111"},{"Street":"Street avenue 2","Country":"Berlin","PostalCode":"90213"}]`

	addresses := &[]Address{}
	err = json.Unmarshal([]byte(jsonAddress), addresses)

	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestJsonDecodeTag(t *testing.T) {
	// reading json not sensitive case
	jsonString 	:= `{"ID":"FE-0001","name":"Flat Enterprise","price":250000,"image_url":"https://google.com"}`
	jsonBytes 	:= []byte(jsonString)

	// type Person in file encode_json_test.go
	product := &Product{}

	err := json.Unmarshal(jsonBytes, product)

	if err != nil {
		panic(err)
	}

	fmt.Println(product.Id)
	fmt.Println(product.Name)
	fmt.Println(product.Price)
	fmt.Println(product.ImageUrl)
}

func TestJsonDecodeMap(t *testing.T) {
	// reading json not sensitive case
	jsonString 	:= `{"id":"FE-0001","ImageUrl":"https://google.com","NAME":"Flat Enterprise","Price":250000}`
	jsonBytes 	:= []byte(jsonString)

	// type Person in file encode_json_test.go
	var product map[string]interface{}

	err := json.Unmarshal(jsonBytes, &product)

	if err != nil {
		panic(err)
	}

	fmt.Println(product)
	fmt.Println(product["id"])
	fmt.Println(product["NAME"])
	fmt.Println(product["Price"])
	fmt.Println(product["ImageUrl"])
}