package routin

import (
	"fmt"
	"runtime"
	"testing"
)

func TestMacProc(t *testing.T) {
	runtime.GOMAXPROCS(12)
	total := runtime.GOMAXPROCS(-1)
	pros := runtime.NumCPU()

	fmt.Println("tread ", total)
	fmt.Println("Prosecors ", pros)
}
