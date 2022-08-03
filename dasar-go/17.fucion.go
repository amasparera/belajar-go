package dasar

func Parameter(nama string) {}

func ReturnData() string {
	return ""
}

func Multiple() (string, string) {
	return "", ""
}

// named parameter
func GetFullName() (name, fulname, last string) {
	return name, "", ""
}

// variadic funcuin
func Variadic(number ...int) int {
	return number[0]
}

// func as parameted
func ParameterFunc(filter func(string) string) {

}

func Tanpanama() {}
