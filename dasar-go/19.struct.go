package dasar

import "fmt"

// comintar
/**
pembuatan struct dasar

**/
func PembuatanStruct() {

	// cara pertama

	var amas User

	amas.Age = 5
	amas.Andres = "joni"
	amas.Name = "julian"

	// cara ke dua
	joko := User{Name: "joni", Andres: "", Age: 5}

	// cara ke tiga
	asu := User{"", "", 7}

	fmt.Println(joko.Age + amas.Age + asu.Age)
	amas.PrintIsi()
}

// sebuat tlamplate datat
// data struct di simpan dalam field

type User struct {
	Name, Andres string
	Age          int
}

// struct method

func (user User) PrintIsi() {
	fmt.Println("ini isi ", user.Name)
}
