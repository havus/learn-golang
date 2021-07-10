package main

import (
	// "os"
	// "flag"
	// "strings"
	// "strconv"
	// "math"
	// "container/list"
	// "container/ring"
	// "sort"
	// "time"
	// "reflect"
	"regexp"
	"fmt"
)

func main() {
	// >>>>>>>>>>> PACKAGE OS
	// $ go run package.go coba argumen
	// args := os.Args
	// fmt.Println(args)

	// fmt.Println(args[1])
	// fmt.Println(args[2])

	// hostname, _ := os.Hostname()
	// fmt.Println(hostname)

	// >>>>>>>>>>> PACKAGE Flag
	// $ go run package.go -host="hello world" -key=root
	// var host *string = flag.String("host", "defaultHost", "Just a description")
	// *key = flag.String("key", "defaultKey", "Insert ur key")
	// // var key *string = flag.String("key", "defaultKey", "Insert ur key")
	// // host := flag.String("hostname", "defaultValue", "Just a description")

	// flag.Parse()

	// fmt.Println(*host)
	// fmt.Println(*key)

	// >>>>>>>>>>> PACKAGE STRING
	// sentence := "Hello World"
	// fmt.Println(strings.Contains(sentence, "hello")) // false
	// fmt.Println(strings.Split(sentence, " "))

	// fmt.Println(strings.ToLower(sentence))
	// fmt.Println(strings.ToUpper(sentence))
	// fmt.Println(strings.ToTitle(sentence))

	// sentence = "   " + sentence + "   "
	// fmt.Println(sentence)
	// fmt.Println(strings.Trim(sentence, " "))

	// fmt.Println(strings.ReplaceAll(sentence, "Hello", "world"))

	// >>>>>>>>>>> STRING CONVERSION
	// bool, err := strconv.ParseBool("true")
	// float, _ := strconv.ParseFloat("3.1415", 64)
	// number, _ := strconv.ParseInt("-42", 10, 64)
	// unassignedInteger, _ := strconv.ParseUint("42", 10, 64)

	// fmt.Println(bool, float, number, unassignedInteger, err)

	// bool := strconv.FormatBool(true)
	// float := strconv.FormatFloat(3.1415, 'E', -1, 64)
	// number := strconv.FormatInt(-42, 10)
	// unassignedInteger := strconv.FormatUint(42, 10)

	// fmt.Println(bool, float, number, unassignedInteger)

	// number, _ := strconv.Atoi("100000")
	// str := strconv.Itoa(100009)

	// fmt.Println(number, str)

	// >>>>>>>>>>> MATH
	// round := math.Round(1.5)
	// ceil 	:= math.Ceil(1.1)
	// floor := math.Floor(1.9)
	// max 	:= math.Max(10, 1)
	// min 	:= math.Min(10, 1)

	// fmt.Println(round, ceil, floor, max, min)

	// >>>>>>>>>>> CONTAINER LIST
	// data := list.New()

	// data.PushBack("World")
	// data.PushBack("John")
	// data.PushBack("Doe")
	// data.PushFront("Hello")

	// fmt.Println(data.Front().Next().Next().Value)
	// fmt.Println(data.Back().Prev().Value)
	// fmt.Println(data.Back().Next())

	// fmt.Println(data)

	// for element := data.Front(); element.Value == "World" || element.Value == "Hello"; element = element.Next() {
	// for element := data.Front(); element != nil; element = element.Next() {
	// 	fmt.Println(element.Value)
	// }
	// for element := data.Back(); element != nil; element = element.Prev() {
	// 	fmt.Println(element.Value)
	// }

	// >>>>>>>>>>> CONTAINER RING
	// data := ring.New(7)

	// for i := 0; i < data.Len(); i++ {
	// 	data.Value = "value ke-" + strconv.FormatInt(int64(i), 16)
	// 	data = data.Next()
	// }

	// data.Do(func(value interface{}) {
	// 	fmt.Println(value)
	// })

	// >>>>>>>>>>> SORT
	// var persons PersonSlice = []Person{
	// 	{"Hello", 2},
	// 	{"World", 12},
	// 	{"John", 1},
	// 	{"Doe", 20},
	// }
	// sort.Sort(PersonSlice(persons))
	// fmt.Println(persons)

	// slice_integer := []int{1, 10, -3, 0, 99, 100}
	// fmt.Println(slice_integer)
	
	// sort.Ints(slice_integer)
	// fmt.Println(slice_integer)

	// >>>>>>>>>>> TIME
	// var currentTime 	= time.Now()
	// var currentDate 	= time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC)
	// var layout 				= time.RFC3339
	// var layout 				= "2006-01-02"
	// var parsedDate, _	= time.Parse(layout, "2020-01-01")

	// fmt.Println(currentTime)
	// fmt.Println(currentDate)
	// fmt.Println(parsedDate)

	// >>>>>>>>>>> REFLECT
	// budi := Person{"Budi", 90}
	// var sampleType reflect.Type 	= reflect.TypeOf(budi)
	// var sampleValue reflect.Value = reflect.ValueOf(budi)

	// fmt.Println(sampleType) // main.Person
	// fmt.Println(sampleType.NumField()) // 2
	// fmt.Println(sampleType.Field(0)) // {Name  string required:"true" 0 [0] false}
	// fmt.Println(sampleType.Field(0).Type) // string
	// fmt.Println(sampleType.Field(0).Tag.Get("required"))
	// fmt.Println(sampleType.Field(1).Tag.Get("max"))

	// fmt.Println(sampleValue.Field(0).Interface())

	// fmt.Println(IsValid(budi))

	// >>>>>>>>>>> REGEX
	regexEmail 		:= "^[a-z0-9._%+\\-]+@[a-z0-9.\\-]+\\.[a-z]{2,4}$"
	regexdowncase	:= "[a-z]*"
	regex 				:= regexp.MustCompile(regexEmail)

	fmt.Println(regex.MatchString("test44$@gmail.com"))
	fmt.Println(regexp.MustCompile(regexdowncase).FindAllString("test hello1@world.com", -1))
}

// func IsValid(data interface{}) bool {
// 	typeOfData := reflect.TypeOf(data)

// 	for i := 0; i < typeOfData.NumField(); i++ {
// 		field := typeOfData.Field(i)

// 		isRequired, _ := strconv.ParseBool(field.Tag.Get("required"))

// 		if isRequired {
// 			value := reflect.ValueOf(data).Field(i).Interface()
// 			return value != ""
// 		}
// 	}
// 	fmt.Println(data)

// 	return true
// }

// type Person struct {
// 	Name 	string 	`required:"true"`
// 	Age 	int			`max:"10"`
// }

// type PersonSlice []Person

// func (object PersonSlice) Len() int {
// 	return len(object)
// }

// func (object PersonSlice) Less(i, j int) bool {
// 	return object[i].Age < object[j].Age
// }

// func (object PersonSlice) Swap(i, j int) {
// 	object[i], object[j] = object[j], object[i]
// }
