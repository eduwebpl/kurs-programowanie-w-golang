package main

import "fmt"

func main() {
	cityTemperature := CityTemperature{[]float64{-1.8, -0.8, 2.9, 8.8, 14.2, 16.8, 19.2, 18.5, 13.8, 8.8, 3.2, -0.7}}
	printCityInfo(cityTemperature)
}

func printCityInfo(city City) {
	fmt.Println(city.averageTemperature())
}

type City interface {
	averageTemperature() float64
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
