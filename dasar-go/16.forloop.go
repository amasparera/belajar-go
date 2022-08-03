package dasar

import "fmt"

func CetakPerulangan() {
	counter := 1
	for counter < 5 {
		fmt.Println(true)
		counter++
	}

	for counter := 1; counter < 5; counter++ {
		fmt.Println(false)

	}
}
