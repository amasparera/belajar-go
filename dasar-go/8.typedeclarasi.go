package dasar

import "fmt"

func CetakTypeDeclarasi() {
	type NoKtp string

	var ktp NoKtp = "98273435740"
	fmt.Println(ktp)
}
