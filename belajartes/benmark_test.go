package belajartes

import (
	"belajar-go/dasar-go"
	"testing"
)

func BenchmarkHelloWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dasar.HelloWord()
	}
}

//  go test -v -bench=. menjalankan seluruh benchmark di module,
//  go test -v -run=NotMathUnitTest -bench=. benchmark tanpa unit test
// 	go test -v -run=NotMathUnitTest -bench=Nama_benchmark menjalankan tertentu
// 	go test -v -bench=../...
