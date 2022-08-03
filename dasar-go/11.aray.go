package dasar

import "fmt"

func CetakAraay() {
	var names [3]string
	names[0] = "eko"

	var name = [3]string{
		"langsung",
		"ok",
		"b aja",
	}

	fmt.Println(names)
	fmt.Println(name)

	fmt.Println(len(names))
	fmt.Println(len(name))

}
