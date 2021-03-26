package main

import (
	"fmt"
	"strconv"
	"reflect"
	"strings"
)

/**
This is about function:
1. Normal function
2. Return value function
3. Return named value
4. Anonymous Function
5. Function as Parameter
6. Recursive Function
7. Defer, Panic & Recover
*/

func main() {
	fmt.Println("================== Function ==================")
	// resultFunction := make([]interface{}, 10)
	stringQuestion, resultFunction := add(1000000, 2)
	stringQuestion1, stringQuestion2 := addWithString(1000000, 2)
	saySomething(stringQuestion + strconv.Itoa(resultFunction))
	saySomething(stringQuestion1 + stringQuestion2)

	printAll([]string{"john", "doe", "hello", "world"}...)

	var coba []interface{}
	coba = append(coba, "hello", "world", 1, 2, 3)
	fmt.Println(coba...)

	saySomethingWithFilter("halo asu anjing babi!", filterKataKasar)

	// Anonymous Function
	saySomethingWithFilter(
		"halo admin!",
		func(words string) string {
			return strings.Replace(words, "admin", "*****", -1)
		},
	)

	fmt.Println(countFactorial(5))

	fmt.Println(divideNumbers(10, 2, 0)) // panic: runtime error: integer divide by zero
	printApapun("Hello,", 1, true)

	// interface return converter
	fmt.Println(countFactorial(returnFive().(int)))
	// fmt.Println(countFactorial(returnFive())) // error
}

// normal function
func saySomething(str string) { // func saySomething(str, str2 string) {}
	fmt.Println(str)
}

// return value function
func add(num1 int, num2 int) (string, int) { // func add(num1 int, num2 int) int {}
	stringQuestion := "hasil dari " + strconv.Itoa(num1)
	stringQuestion += " + " + strconv.Itoa(num2) + " adalah "

	fmt.Println(reflect.ValueOf(num1).Kind(), reflect.ValueOf(num2).Kind())

	return stringQuestion, num1 + num2
}

// return named value
func addWithString(num1 int, num2 int) (stringQuestion, answer string) { // func add(num1 int, num2 int) (stringQuestion string, answer int) {}
	stringQuestion = "hasil dari " + strconv.Itoa(num1)
	stringQuestion += " + " + strconv.Itoa(num2) + " adalah "

	answer = strconv.Itoa(num1 + num2)

	return
}

func printAll(dummyString ...string) {
	for _, value := range dummyString {
		fmt.Print(value, ", ")
	}
	fmt.Println("\nThat's Variadic Function")
}

func filterKataKasar(word string) string {
	// With strings.Replace we have a way to target one or more instances of substrings.

	//  The final argument is the occurrence count to replace.
	// We pass -1 as the occurrence count, so all instances are replaced.
	// return strings.Replace(word, "asu", "***", -1)

	// With Replacer, we have a powerful class that can combine many replacements.
	// This will speed up large or complex sets of replacement.
	rant := []string{
		"asu", 		"***",
		"anjing",	"*****",
		"babi", 	"****",
	}
	replacer := strings.NewReplacer(rant...)

	return replacer.Replace(word)
}

// Function as Parameter
type Filter func(string) string

func saySomethingWithFilter(words string, filter Filter) {
	fmt.Println(filter(words))
}

// Recursive Function
func countFactorial(value int) int {
	if value == 1 {
		return 1
	}

	return value * countFactorial(value - 1)
}

// Defer, Panic & Recover
func divideNumbers(numbers ...int) int {
	defer endFunction()
	value := numbers[0]

	// panic("boom!")
	for _, number := range numbers[1:] {
		value /= number
	}

	return value
}

func endFunction() {
	errorMessage := recover()
	fmt.Println("Error message adalah", errorMessage)
	fmt.Println("Function selesai!")
}

func printApapun(params ...interface{}) {
	for _, param := range params {
		fmt.Println(param)
	}
}

func returnFive() interface{} {
	return 5
}
