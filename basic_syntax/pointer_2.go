package main

import "fmt"

// type error interface {
// 	Error() string
// }

type Person struct {
	Name  string
  Age   int
}

// NOT CHANGE
// func (person Person) ChangeName(newName string) {
//   person.Name = newName
// }

// NOT CHANGE
func (person *Person) ChangeName(newName string) {
  person.Name = newName
}

func (person *Person) ChangeAge(newAge int) {
  person.Age = newAge
}

// func (error *AppError) AsMessage() AppError {
//   return AppError{ Message: e.Message }
// }

// func (error AppError) AsMessage() *AppError {
//   return &AppError{ Message: e.Message }
// }

// func (error *AppError) AsMessage() *AppError {
//   return &AppError{ Message: e.Message }
// }

func duplicatePerson(name *string, age *int) *Person {
  return &Person{ Name: *name, Age: *age }
}

func changeInt(number *int) {
  *number = 100
}

func main() {
  budi    := Person{ Name: "Budi", Age: 20 }
  newAge  := 30

  // fmt.Println(&newAge) // 0xc000014098
  // fmt.Println(*&newAge) // 30
  changeInt(&newAge)

  budi.ChangeName("Ganti Budi")
  budi.ChangeAge(newAge)

  fmt.Println(budi)

  joko := duplicatePerson("joko")
  fmt.Println(joko)
}
