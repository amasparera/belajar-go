package pakege

import (
	"fmt"
	"os"
)

func CetasOs() {
	dat := os.Args
	hosnam, _ := os.Hostname()

	pssword := os.Getenv("APP_Database")

	fmt.Println(dat)
	fmt.Println(hosnam)
	fmt.Println(pssword)
}

// terminal
// export APP_Databaase=root
// export APP_Password=root
