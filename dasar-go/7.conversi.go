package dasar

import (
	"fmt"
)

func Cetakkonversi() {

	var n32 int32 = 33333
	var n64 int64 = int64(n32)

	fmt.Println(n64)

	var n8 int8 = int8(n64)

	fmt.Println(n8)

	var name = "eko"
	var e = name[0]
	fmt.Println(e)

	fmt.Println(string(name[1]))

	var s = string(e)
	fmt.Println(s)

}
