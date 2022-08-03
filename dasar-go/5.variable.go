package dasar

import "fmt"

// single
var name string

// multi
var (
	asu   = "ajing"
	wedos = "kambing"
)

func CetakVariable() {
	name = "amas parere"
	fmt.Println(name)

	name = "nama baru"
	fmt.Println(name)

	fmt.Println(asu + wedos)
}
