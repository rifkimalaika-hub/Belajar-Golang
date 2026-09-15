package main

import "fmt"

func main() {

	// Array
	names := [4]string{
		2: "Nurmanto",
		0: "Rifki",
		3: "S.kom",
		1: "Malaika",
	}
	
	fmt.Println(names[0])
	fmt.Println(names)

	// Array dengan panjang yang tidak ditentukan
	values := [...]int{
		4: 10,
		2: 20,
		1: 30,
	}
	fmt.Println(values)
	fmt.Println(len(values))

}	