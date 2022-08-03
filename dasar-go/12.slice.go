package dasar

import "fmt"

// slice data yg mengakses sebgian array

// tipe data slice yaitu pointer, lengthdan capacity

// array [low:hig h dari low sampai high

// len pangjang
// cap capasitas
// appaned(slice, data) membuat slice baru menambah data ke posisi terakhir, jika capasitas penuh membuat array baru
// make ([]typedata , lenght, capasiti) membuat slice baru
// copy (destination, source) mennyalin slice dari destination

func CetakSlice() {
	var moths = [...]int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12,
	}

	var newSlice = moths[8:11]
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	moths[9] = 0

	// merubah array
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	// merubah slice
	newSlice[0] = 0
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

}
