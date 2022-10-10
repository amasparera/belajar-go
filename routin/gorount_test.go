package routin

import (
	"flag"
	"fmt"
	"testing"
	"time"
)

func Helo() {
	fmt.Println("Hello Word")
}

func NumberAll(index int) {
	fmt.Println("index ke ", index)
}

// bukan untuk mengembalikan value
func TestHelo(t *testing.T) {
	flag.Parse()
	go Helo()
	fmt.Println("ups")

	time.Sleep(1 * time.Second)

}

func TestDisplay(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go NumberAll(i)
	}
	time.Sleep(5 * time.Second)
}
