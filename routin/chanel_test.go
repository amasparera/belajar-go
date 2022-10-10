package routin

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

// chanel
// komunikasi secara singkronus
// -ada pengirim dan penerima
// -harus close

func TestCreate(t *testing.T) {
	chanel := make(chan string)

	// chanel <- "amas"
	// data1 := <-chanel
	// fmt.Println(data1)

	go func() {
		time.Sleep(2 * time.Second)
		chanel <- "bintang parera"
		fmt.Println("selesai mengirim data")
	}()

	data := <-chanel
	fmt.Println(data)

	time.Sleep(4 * time.Second)

	close(chanel)

}

func GiveMe(chanel chan string) {
	time.Sleep(2 * time.Second)
	chanel <- "masuk pak eko"
}

func TestMasuk(t *testing.T) {
	chanel := make(chan string)

	go GiveMe(chanel)

	data := <-chanel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
	close(chanel)
}

// chanel in dan out

func OnlyIn(chanel chan<- string) {
}

func Onlyout(chanel <-chan string) {}

// buffer chanel
func TestBuffer(t *testing.T) {
	channel := make(chan string, 6)

	channel <- "1"
	channel <- "6"

	fmt.Println(len(channel))
	fmt.Println((<-channel))
	fmt.Println((<-channel))

	close(channel)
}

// range channel

func TestRangeChannel(t *testing.T) {
	chanel := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			chanel <- "perulangan " + strconv.Itoa(i)
			fmt.Println("masuk ", i)
		}
		close(chanel)
	}()

	for data := range chanel {
		fmt.Println(data, " diterima")
	}

}

// select chanel

// log chanel kunci race kondition
func TestLog(t *testing.T) {
	var x = 0
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()

	}

	time.Sleep(5 * time.Second)
	fmt.Println("jumlah ", x)
}
