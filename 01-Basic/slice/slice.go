package main

import "fmt"

func main() {
	// names := [...]string{"Rifki", "Malaika", "Nurmanto", "S.kom"}

	// names1 := names[1:3] // mengambil data dari index 1 sampai index 2 (index 3 tidak termasuk)
	// fmt.Println(names1)

	// names2 := names[2:4] // mengambil data dari index 2 sampai index 3 (index 4 tidak termasuk)
	// fmt.Println(names2)
	
	// names3 := names[:2] // mengambil data dari index 0 sampai index 1 (index 2 tidak termasuk)
	// fmt.Println(names3)
	
	// names4 := names[:] // mengambil semua data dari index 0 sampai index terakhir
	// fmt.Println(names4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	fmt.Println(days)

	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu baru"
	daySlice1[1] = "Minggu baru"
	fmt.Println(daySlice1)
	fmt.Println(days) // data pada array days ikut berubah karena slice daySlice1 mereferensikan data pada array days

	// Slice Append
	daySlice2 := append(daySlice1, "Libur") // menambahkan data baru ke dalam slice daySlice1, sehingga membuat array baru [Sabtu baru, Minggu baru, Libur]
	daySlice2[0] = "Sabtu lagi" // [Sabtu lagi, Minggu baru, Libur]
	fmt.Println(daySlice2)
	fmt.Println(days) // data pada array days tidak ikut berubah karena slice daySlice2 tidak mereferensikan data pada array days, melainkan membuat array baru

	// Make Slice
	newSlice := make([]string, 3, 5) // membuat slice baru dengan panjang 3 dan kapasitas 5
	newSlice[0] = "Rifki"
	newSlice[1] = "Malaika"
	newSlice[2] = "Nurmanto"
	// newSlice[3] = "S.kom" // akan terjadi panic karena panjang slice hanya 3, sehingga index 3 tidak bisa diakses

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "S.kom") // [Rifki, Malaika, Nurmanto, S.kom] menambahkan data baru ke dalam slice newSlice, sehingga membuat array baru
	newSlice2[1] = "Dera"
	fmt.Println(newSlice2)
	fmt.Println(newSlice) // data pada slice newSlice tidak ikut berubah karena slice newSlice2 tidak mereferensikan data pada slice newSlice, melainkan membuat array baru

	// Copy Slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(toSlice)
	fmt.Println(fromSlice)

}