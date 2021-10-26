package main

import "fmt"

// uruchom za pomocą komendy
// go run ./main-zadanie-finish.go

func main() {
	warsaw := CityTemperature{[]float64{-1.8, -0.8, 2.9, 8.8, 14.2, 16.8, 19.2, 18.5, 13.8, 8.8, 3.2, -0.7}}
	madrid := CityTemperature{[]float64{6, 7.6, 10.8, 12.6, 16.5, 22.2, 25.5, 25.2, 21, 15.2, 9.9, 6.7}}
	fmt.Println(warsaw.averageTemperature())
	fmt.Println(madrid.averageTemperature())
}

type CityTemperature struct {
	temperatures []float64
}

func (c CityTemperature) averageTemperature() float64 {
	var sumOfTemps float64 = 0
	for _, temperature := range c.temperatures {
		sumOfTemps += temperature
	}
	return sumOfTemps / float64(len(c.temperatures))
}
