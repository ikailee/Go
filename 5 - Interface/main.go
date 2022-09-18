package main

import "fmt"

type Printer interface {
	printInfo()
}

type Book struct {
	Title string
	Price float32
}

type Drink struct {
	Name  string
	Price float32
}

func (b Book) printInfo() {
	fmt.Printf("%s, %.2f", b.Title, b.Price)
}

func (d Drink) printInfo() {
	fmt.Printf("%s, %.2f", d.Name, d.Price)
}

func empty(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("string: %s", i)
	case int:
		fmt.Printf("string: %s", i)
	case Book:
		fmt.Printf("string: %s", i.(Book).Title)
	}

}

func main() {
	book := Book{
		Title: "Book",
		Price: 9.99,
	}
	book.printInfo()

	drink := Drink{
		Name:  "Drink",
		Price: 5.99,
	}
	drink.printInfo()

	info := []Printer{book, drink}
	info[0].printInfo()
	info[1].printInfo()

	empty("kevin")
	empty(100)
	empty(book)
}
