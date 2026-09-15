package main

import "fmt"

func main() {
	
	// // Operator Perbandingan
	// k := 10
	
	// nama1 := "Rifki"
	// nama2:= "Malaika"
	
	// fmt.Println("Apakah k lebih besar dari 5 =", k > 5)
	// fmt.Println("Apakah k lebih kecil dari 5 =", k < 5)
	// fmt.Println("Apakah k sama dengan 10 =", k == 10)
	// fmt.Println("Apakah k tidak sama dengan 10 =", k != 10)
	
	// fmt.Println("Apakah nama1 sama dengan nama2 =", nama1 == nama2)
	// fmt.Println("Apakah nama1 tidak sama dengan nama2 =", nama1 != nama2)
	
	// Operator Logika
	nilaiAkhir := 80
	Absensi := 70

	lulusNilai := nilaiAkhir > 85
	lulusAbsensi := Absensi > 60
	lulus := lulusNilai && lulusAbsensi
	lulus2 := lulusNilai || lulusAbsensi
	fmt.Println("Apakah lulus berdasarkan nilai?", lulusNilai)
	fmt.Println("Apakah lulus berdasarkan absensi?", lulusAbsensi)
	fmt.Println("Apakah lulus berdasarkan kedua kondisi?", lulus)
	fmt.Println("Apakah lulus berdasarkan salah satu kondisi?", lulus2)
}