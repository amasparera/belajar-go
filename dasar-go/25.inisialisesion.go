package dasar

import "fmt"

var Data string

// initstate
func init() {
	Data = "init state"

}

func CetakData() {
	fmt.Println(Data)
}
