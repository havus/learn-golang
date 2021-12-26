package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var url string = "https://api.github.com"

func main()  {
	// firstWay()
	secondWay()
}

func jsonPrint(jsonBytes []byte) {
	var prettyJson bytes.Buffer

	// json.Indent(&prettyJson, jsonBytes, "", "\t")
	// \t is a symbol for tab
	if err := json.Indent(&prettyJson, jsonBytes, "", "   "); err != nil {
		panic(err)
	}

	fmt.Println(string(prettyJson.Bytes()))
}

func firstWay() {
	client := http.Client{}

	res, err 	:= client.Get(url)
	resBody 	:= res.Body

	defer resBody.Close()

	if err != nil {
		panic(err)
	}

	bs, err := ioutil.ReadAll(resBody)
	if err != nil {
		panic(err)
	}

	fmt.Println("status:", res.StatusCode)
	jsonPrint(bs)
}

func secondWay() {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("Accept", "application/json")

	client 		:= http.Client{}
	res, err 	:= client.Do(request)
	resBody 	:= res.Body

	defer resBody.Close()

	if err != nil {
		panic(err)
	}

	bs, err := ioutil.ReadAll(resBody)
	if err != nil {
		panic(err)
	}

	fmt.Println("status:", res.StatusCode)
	jsonPrint(bs)
}