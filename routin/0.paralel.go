package routin

// parallel
// memecah masalah menjadi lebih kecil dan di jalan kan secara bersamaan
// -banyak tread

// cocarensi
// mengerjakan menjalan kan pekerjaan secara bergantian
// -sedikit tread

// proses = exsekusi program , mengosumsi memory, saling ter isolali, lama jalan
// tread = treand pecahan proses menjadi kecil

// cpu bound keceaptan cpu
// I/O bound data dari file

// goroutine tread ringan di kelola go runtime
// -berjalan curanren di dalam tread
// -go schedule
// -jumlah GomaxProscs
// - g: gorountine m: tread p: procesor
