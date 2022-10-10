package belajartes

import (
	"belajar-go/dasar-go"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWord(t *testing.T) {
	dasar.HelloWord()

	// Fail() menggagalkan unit test dan melanjutkannya code berikutnya, tetep fuction selanjutnya
	// Error(Arg..) dengan print
	// FailNow() menggagalkan unit test dan berhenti ,tetep fuction selanjutnya
	// Fatal(Arg..) dengan print berhenti

}

// testing.T untuk testing
// testing.M untuk mengatur life cycle testing
// testing.B untuk melkukan benchmarking
// go test untuk menjalankan
// go test -v untuk run dan melihat file apa saja yg jalan
// go test -v -run=nama_fuction untuk menjalankan satu fuction
// go test -v -run nama_file untuk menjalankan satu file
// go test./... menjalankan semua

// assertion  https://github.com/stretchr/testify
// require perbedaan fail dan failnow

func TestAsse(t *testing.T) {
	result := dasar.HelloWord()

	assert.Equal(t, "Hallo amas", result)
}

// skip test t.Skip(arg..) mebatalkan fuction test

// sub tes di dalam fuction t.Run("arg..",func)
