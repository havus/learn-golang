package learn_golang_json

import (
	"encoding/json"
	"testing"
	"fmt"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJsonEncodeDataType(t *testing.T) {
	logJson("Hello world")
	logJson(3)
	logJson(true)
	logJson([]string{"hello", "world"})
}
type Address struct {
	Street			string
	Country			string
	PostalCode	string
}

type Person struct {
	FirstName		string
	LastName		string
	Age					int
	IsStudent		bool
	Hobbies			[]string
	Addresses		[]Address
}

type Product struct {
	Id 				string 	`json:"id"`
	Name 			string 	`json:"name"`
	Price 		int 		`json:"price"`
	ImageUrl 	string 	`json:"image_url"`
}

func TestJsonEncode(t *testing.T) {
	johnDoe := Person{
		FirstName: 	"John",
		LastName: 	"Doe",
		Age: 				20,
		Hobbies:		[]string{"coding", "reading"},
		Addresses:	[]Address{
			{
				Street: "Street avenue",
				Country: "Swiss",
				PostalCode: "72111",
			},
			{
				Street: "Street avenue 2",
				Country: "Berlin",
				PostalCode: "90213",
			},
		},
	}

	bytes, _ := json.Marshal(johnDoe)
	fmt.Println(string(bytes))

	fmt.Println("============================== JSON ARRAY ==========================")

	addresses := []Address{
		{
			Street: "Street avenue",
			Country: "Swiss",
			PostalCode: "72111",
		},
		{
			Street: "Street avenue 2",
			Country: "Berlin",
			PostalCode: "90213",
		},
	}

	addressBytes, _ := json.Marshal(addresses)

	fmt.Println(string(addressBytes))
}

func TestJsonTag(t *testing.T) {
	product := Product{
		Id: 				"FE-0001",
		Name: 			"Flat Enterprise",
		Price: 			250000,
		ImageUrl: 	"https://google.com",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJsonMap(t *testing.T) {
	// can NOT use tag while enncode
	product := map[string]interface{}{
		"id": 			"FE-0001",
		"NAME": 		"Flat Enterprise",
		"Price":		250000,
		"ImageUrl":	"https://google.com",
	}

	// bytes, _ := json.Marshal(product)
	bytes, _ := json.MarshalIndent(product, "", "    ") // pretty-print
	fmt.Println(string(bytes))
}
