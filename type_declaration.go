package main

import "fmt"

type NoKTP string //NoKTP adalah type data string
type Married bool //Married adalah type data boolean
type Stock int //stock adalah type data integer

func main() {
	// contoh type conversion string
	KTPrifki := NoKTP("112222222")
	test := "121299992"
	testKTP := NoKTP(test)
	fmt.Println(KTPrifki)
	fmt.Println(testKTP)

	// contoh type conversion boolean
	status := Married(true)
	status2 := bool(status)
	fmt.Println(status)
	fmt.Println(status2)

	// contoh type conversion integer
	stockA := Stock(100)
	stockB := int(stockA)
	fmt.Println(stockA)
	fmt.Println(stockB)
}