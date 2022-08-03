package dasar

import "fmt"

// mengubah tipe data menjani tipe data yg diingin kan
// serindigunakan ketika data interfice kosong

func Random() interface{} {
	return true
}

func CetakRandom() {
	result := Random()
	switch value := result.(type) {
	case string:
		fmt.Println("value ", value, " is string")
	case int:
		fmt.Println("value ", value, " is int")
	default:
		fmt.Println("Unknown type")
	}

}
