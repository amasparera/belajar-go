package dasar

import (
	"fmt"
)

// defer funcion adalah jadwal exsekusi bila selesai
// selalu di eksekusi walaupun function err

func DeferHandle() {
	defer HelloWord()
}

// panic
// menghentikan program
// dipanggib ketika program kita error
// function berhenti tapi defer berjalan

func PanicHandle(input uint8) {
	if input == 0 {
		panic("tidak berjalan")
	}

}

// recaver
// menangkap data panic
// dengan recaveri proses panic berhenti, program tetap berjalan

func RecaverU() {
	message := recover()
	fmt.Println(message)

	PanicHandle(0)
}
