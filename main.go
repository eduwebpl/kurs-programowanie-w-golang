package main

import (
	"fmt"
	"runtime"
)

const i int = 0

func main() {
	// var i int = 0
	// var i = 0
	// i := 0

	// var tablica []int = []int{5, 3, 4}
	// tablica := []int{5, 3, 4}

	// tablica[0] // - 5
	// tablica[2] // - 4

	// len(tablica) // - 3

	exampleMap := map[string]string{}
	exampleMap["key"] = "value"
	fmt.Println(exampleMap["key"])

	if true == true {

	}

	if 1 == 2 { // jeśli to

	} else if 3 == 4 { // jeśli nie tamto to wtedy jeśli to

	} else { // a jeśli nie cała reszta to wtedy to

	}

	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println(os)
	}

	i := 0
	for i < 10 {
		i += 1
	}

	shouldEnd := false

	for shouldEnd == false {

	}

	tablica := []int{5, 3, 4}
	for index, element := range tablica {
		fmt.Println(index)
		fmt.Println(element)
	}

	tablica = []int{5, 3, 4}
	for _, element := range tablica {
		fmt.Println(element)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}

func createHTTPServer() {

}

// Zadanie domowe: Policz średnią temperaturę.

// Warszawa:
// -1.8, -0.8, 2.9, 8.8, 14.2, 16.8, 19.2, 18.5, 13.8, 8.8, 3.2, -0.7

// Madryt:
// 6, 7.6, 10.8, 12.6, 16.5, 22.2, 25.5, 25.2, 21, 15.2, 9.9, 6.7
