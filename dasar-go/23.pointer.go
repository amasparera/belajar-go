package dasar

import "fmt"

// pasing by value
type Adress struct {
	cty, name string
}

func Reference() {
	adress1 := Adress{"jombang", "jatim"}
	// dikopy dan 2 nilai berbeda
	adress2 := adress1

	adress2.cty = "surabaya"

	fmt.Println(adress1)
	fmt.Println(adress2)
}

func ReferenceBypointer() {
	adress1 := Adress{"jombang", "jatim"}
	// pointer dan referen value nya
	adress2 := &adress1

	adress2.cty = "surabaya"

	fmt.Println(adress1)
	fmt.Println(adress2)
}

var Adress4 *Adress = new(Adress)

// parameter pointer
func FuncParameter() {
	Adress4 = &Adress{"jogja", "indonesia"}

	fmt.Println(Adress4)
}

// pointer tipe pointer ke memori berbeda
var Andres5 Adress = Adress{}
var Adress6 = &Andres5

func MerubahSatu() {
	Adress6 = &Adress{"mediun", "ok"}

	fmt.Println(Andres5)
	fmt.Println(Adress6)
}

// merubah pointer kebaru semua yg mengacu
func MerubahKeMemoriBaru() {
	*Adress6 = Adress{"kota", "kita"}

	fmt.Println(Andres5)
	fmt.Println(Adress6)
}
