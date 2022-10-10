package routin

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Bank struct {
	RwMutex sync.RWMutex
	Uang    int
}

func (bank *Bank) TambahNominal(nominal int) {
	bank.RwMutex.Lock()

	bank.Uang = bank.Uang + nominal
	bank.RwMutex.Unlock()
}

func (bank *Bank) GetNominal() int {
	bank.RwMutex.RLock()
	balace := bank.Uang
	bank.RwMutex.RUnlock()
	return balace
}

// pembuatan test
func TestReadMutext(t *testing.T) {
	nominal := Bank{}

	for i := 0; i < 100; i++ {
		go func() {
			for a := 0; a < 100; a++ {
				nominal.TambahNominal(1)
				fmt.Println(nominal.GetNominal())
			}
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("final uang : ", nominal.GetNominal())
}

// wait grup
func RunAsycronus(grup *sync.WaitGroup) {

	grup.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
	defer grup.Done()
}

func TestWaitGup(t *testing.T) {
	grup := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsycronus(grup)
	}

	grup.Wait()
	// fmt.Println("Complite")
}

// syn once
// dijalankan func 1x
