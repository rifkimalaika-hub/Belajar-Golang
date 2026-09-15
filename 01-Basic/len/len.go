package main

import "fmt"

func main() {

	a := [5]int{1, 2, 3, 4, 5}
	result := a[2:3]
	fmt.Println(a) // [1 2 3 4 5]
	fmt.Println(result) // mengambil data dari index 2 sampai index 2 (index 3 tidak termasuk)
	fmt.Println(len(result)) // panjang dari result adalah 1 karena hanya ada 1 data yang diambil yaitu data pada index 2

	b := [4]string{"Rifki", "Malaika", "Nurmanto", "S.kom"}
	fmt.Println(b) // [Rifki Malaika Nurmanto S.kom]
	fmt.Println(len(b)) // panjang dari b adalah 4 karena ada 4 data di dalam array b
}
