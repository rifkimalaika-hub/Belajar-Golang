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
	days1 := days[5:]
	days1[0] = "Sabtu baru"
	days1[1] = "Minggu baru"
	fmt.Println(days)

	days2 := append(days1, "Libur")
	days2[0] = "Sabtu lagi"
	fmt.Println(days2)
	fmt.Println(days)


}