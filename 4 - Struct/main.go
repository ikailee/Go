package main

import (
	"encoding/json"
	"fmt"
)

type Rectangle struct {
	Length int
	Width  int
}

// value receiver
func (r Rectangle) getArea() int {
	return r.Length * r.Width
}

// pointer receiver
func (r *Rectangle) setLength(len int) {
	r.Length = len
}

type Person struct {
	Name string `json:"name"`
	Age  uint8
}

type Employee struct {
	Id     uint64
	Person Person
	Salary uint
}

func main() {
	var rect1 Rectangle
	rect1.Length = 1
	rect1.Width = 2
	fmt.Printf("%v", rect1.getArea())

	rect2 := Rectangle{}
	rect2.Length = 3
	rect2.Width = 4

	rect3 := Rectangle{5, 6}
	fmt.Printf("%v, %v", rect3.Length, rect3.Width)

	rect4 := Rectangle{
		Length: 7,
		Width:  8,
	}
	fmt.Printf("%v, %v", rect4.Length, rect4.Width)

	var rect5 = new(Rectangle)
	rect5.Length = 9
	rect5.Width = 10

	var rect6 = &Rectangle{}
	rect6.Length = 11
	rect6.Width = 12

	kevin := Employee{
		Id:     001,
		Salary: 100000,
		Person: Person{
			Name: "Kevin",
			Age:  40,
		},
	}
	fmt.Printf("%v, %v", kevin.Id, kevin.Person.Name)

	david := Person{
		Name: "David",
		Age:  39,
	}
	data, err := json.Marshal(david)
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Printf("json: %v", data)
}
