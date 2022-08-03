package dasar

import "fmt"

func CetakSwith() {
	name := "joni"

	switch name {
	case "eko":
		fmt.Println("joko")
	case "joni":
		fmt.Println("joni")
	default:
		fmt.Println("jaguar")
	}
}

func ShotExpresion() {
	switch lenght := 5; lenght > 2 {
	case true:
		fmt.Println("true")
	default:
		fmt.Println("false")
	}
}
