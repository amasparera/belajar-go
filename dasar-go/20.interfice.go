package dasar

import "fmt"

// declarasi interface
type HasName interface {
	GetName() string
}

// func SayHello
func SayHello(hasname HasName) {
	fmt.Println("hello", hasname.GetName())
}

// mthod struct
func (person Person) GetName() string {
	return person.Name
}

// struct
type Person struct {
	Name string
}

func CetakInterface() {
	var person Person
	person.Name = "budi"

	SayHello(person)
}

// ini interface Kosong
func Ups() interface{} {
	return "ups"
}
