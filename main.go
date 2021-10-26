package main

func main() {
	// age := 21
	// age2 := age
	// age2 = 22

	age := 21
	age2 := &age
	*age2 = 22

}
