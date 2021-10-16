package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
)

type MarshalPerson struct {
	Name string `json:"name"`
}

func main() {
	andi := MarshalPerson{
		Name: "andi",
	}
	budi := MarshalPerson{
		Name: "budi",
	}

	people 					:= []MarshalPerson{andi, budi}
	dataBytes, err 	:= json.Marshal(people)

	if err != nil {
		fmt.Printf("Error is %s", err)
	}

	fmt.Println(dataBytes)
	fmt.Println(string(dataBytes))
	fmt.Println(reflect.ValueOf(dataBytes).Kind())

	people2 := []MarshalPerson{}

	// data 			:= `[{"name":"andi"},{"name":"budi"}]`
	// dataBytes 	:= []byte(data)
	if err := json.Unmarshal(dataBytes, &people2); err != nil {
		fmt.Printf("Error is %s", err)
	}

	fmt.Println("people2", people2)

	http.HandleFunc("/foo", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)
}

// encode
func foo(w http.ResponseWriter, r *http.Request) {
	andi := MarshalPerson{
		Name: "andi",
	}

	if err := json.NewEncoder(w).Encode(andi); err != nil {
		log.Println("Encoded bad data:", err)
	}
}

// decode
func bar(w http.ResponseWriter, r *http.Request) {
	// > curl --header "Content-Type: application/json" \
	// 	--request POST \
	// 	--data '{"name":"anonim"}' \
	// 	http://localhost:8080/bar

	log.Println("request body", r.Body, "with type", reflect.TypeOf(r.Body))
	// request body &{0xc0000a6288 <nil> <nil> false true {0 0} false false false 0x11faf00} with type *http.body

	log.Println("request body read", r.Body.Read, "with type", reflect.TypeOf(r.Body.Read))
	// request body read 0x1201360 with type func([]uint8) (int, error)

	var person MarshalPerson

	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		log.Println("Encoded bad data: ", err)
	}

	log.Println("Person:", person.Name)
}