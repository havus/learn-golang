package main

import (
	"fmt"
	"reflect"
)

func main() {
  fmt.Println("hello world, this is golang!")

	// >>>>> Type declaration
	type angkaKu int8
	// angka := 80 => error
	var angka angkaKu = 80
	fmt.Println(angka)

	// >>>>> Augmented assignment
	price := 10000
	price += 200
	fmt.Println(price)

	// Unary operator
	price++
	fmt.Println(price)
	isTrue := true
	isFalse := !isTrue
	fmt.Println(isFalse)

	fmt.Println("================== Array ==================")
	var names [2]string

	names[0] = "John"
	names[1] = "Doe"

	fmt.Println(names)

	prices := [3]int{ 10000, 20000, 50000 }
	fmt.Println(prices)

	myDatas := [5]int{ 1, 2, 3 }
	fmt.Println("myDatas length", len(myDatas)) // 5

	myDatas2 := [...]int{ 1, 2, 3 }
	fmt.Println("myDatas2 length", len(myDatas2)) // 3

	fmt.Println("====================================")

	// Slice
	days := [...]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}
	// weekend := days[5:7]
	weekend := days[5:]
	weekdays := days[:6]

	fmt.Println("weekdays", weekdays)
	fmt.Println("weekdays length", len(weekdays))
	fmt.Println("weekdays capacity", cap(weekdays))

	fmt.Println("weekend", weekend)
	fmt.Println("weekend length", len(weekend))
	fmt.Println("weekend capacity", cap(weekend))

	weekend[0] = "Sabtu" // Data reference
	weekend[1] = "Minggu" // Data reference
	fmt.Println("days reference", days)
	fmt.Println("weekend", weekend)

	indoWeekend := append(weekend, "Tanggal Merah") // create new reference
	indoWeekend[0] = "Saturday" // Data NOT reference
	indoWeekend[1] = "Sunday" // Data NOT reference
	fmt.Println("days", days)
	fmt.Println("weekend", weekend)
	fmt.Println("indoWeekend", indoWeekend)

	fmt.Println("================== New Slice ==================")

	newSlice := make([]string, 3, 3) // length, capacity
	newSlice[0] = "Januari"
	newSlice[1] = "Februari"
	newSlice[2] = "March"

	fmt.Println("newSlice", newSlice)
	fmt.Println("newSlice length", len(newSlice))
	fmt.Println("newSlice capacity", cap(newSlice))

	fmt.Println("================== Copy Slice ==================")
	newSlice2 := make([]string, len(newSlice) - 1, cap(newSlice))
	copy(newSlice2, newSlice)

	fmt.Println("newSlice2", newSlice2)
	fmt.Println("newSlice2 length", len(newSlice2))
	fmt.Println("newSlice2 capacity", cap(newSlice2))

	// conclusion
	thisIsArray := [...]string{ "one", "two" }
	thisIsSlice := []string{ "one", "two" }
	fmt.Println("thisIsArray", thisIsArray)
	fmt.Println("thisIsSlice", thisIsSlice)

	fmt.Println("================== Map ==================")
	// make(map[string]string)
	// map[string]string{} // make(map[string]string)
	myMap := map[string]string{ // map[key_data_type]value_data_type
		"first-name":	"John",
		"last-name": 	"Doe",
		"age": 				"18",
	}
	myMap["address"] = "Indonesia"

	fmt.Println("myMap", myMap)
	delete(myMap, "address")
	fmt.Println("myMap", myMap)

	fmt.Println("================== If Statement ==================")

	role := "manager"

	if role == "superadmin" {
		fmt.Println("This is", role)
	} else if role == "manager" {
		fmt.Println("Wow, this is", role)
	} else if roleLength := len(role); roleLength > 10 { // Short statement
		fmt.Println("Role unpredictable!!")
	} else {
		fmt.Println("This is normal users")
	}
	// fmt.Println(roleLength) // will be undefined value

	fmt.Println("================== Switch Statement ==================")

	switch role {
	case "superadmin":
		fmt.Println("This is", role)
	case "manager":
		fmt.Println("Wow, this is", role)
	default:
		fmt.Println("This is normal users")
	}

	switch roleLength := len(role); roleLength { // Short statement
	case 7:
		fmt.Println("This is", role)
	default:
		fmt.Println("This is normal users")
	}
	// fmt.Println(roleLength) // will be undefined value

	role = "role2"
	roleLength := len(role)

	switch {
	case roleLength > 5:
		fmt.Println("lebih dari 5")
	case roleLength > 10:
		fmt.Println("lebih dari 10")
	default:
		fmt.Println("Kurang dari 5")
	}

	fmt.Println("================== Loop ==================")

	counter1 := 0

	for counter1 < 3 {
		fmt.Print(counter1)
		counter1++
	}
	
	for counter := 0; counter < 3; counter++ {
		fmt.Print(counter)
	}
	fmt.Println(" ")

	slice_loop := []string{"hello", "world"}
	fmt.Println(reflect.ValueOf(slice_loop).Kind())

	for i := 0; i < len(slice_loop); i++ {
		fmt.Println(slice_loop[i])
	}

	for i, value := range slice_loop { // for _, value := range slice_loop
		fmt.Println(value, i)
	}

	for key, value := range myMap {
		fmt.Println(key, "=>", value)
	}
}
