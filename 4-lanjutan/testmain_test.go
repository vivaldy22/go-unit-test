package lanjutan_test

import (
	"fmt"
	"os"
	"testing"
)

//TestMain gunanya kalau kita ada variable global yang mau di inisialisasi misalnya, itu bisa disini, tapi inget
//semua yang ada di dalam ini cuman di eksekusi sekali, jadi kalau butuh koneksi db berkali", atau ngisi sebuah
//variabel berkali", harus manual
func TestMain(m *testing.M){
	//before
	fmt.Println("before")
	exit := m.Run()
	//after
	fmt.Println("after")

	os.Exit(exit)
}