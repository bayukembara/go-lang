package main

import "fmt"

func main() {
	var benar bool = true
	var salah bool = false

	fmt.Println(benar && benar)
	fmt.Println(benar && salah)
	fmt.Println(salah && benar)
	fmt.Println(salah && salah)

	nilaiAkhir := 90
	absensi := 90

	lulusNilaiAkhir := nilaiAkhir > 80
	lulusAbsensi := absensi > 80

	lulus := lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)

}
