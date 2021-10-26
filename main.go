package main

import (
	"fmt"
	"sort"
)

func main() {
	cityTemperature := CityTemperature{[]float64{-1.8, -0.8, 2.9, 8.8, 14.2, 16.8, 19.2, 18.5, 13.8, 8.8, 3.2, -0.7}}
	fmt.Println(cityTemperature.temperatures)
	cityTemperature.sort()
	fmt.Println(cityTemperature.temperatures)

	// Map
	newTemperatures := cityTemperature.toNewCityTemperatures()

	// Filter
	temperaturesAboveZero := []float64{}
	for _, temperature := range cityTemperature.temperatures {
		if temperature > 0 {
			temperaturesAboveZero = append(temperaturesAboveZero, temperature)
		}
	}
	fmt.Println(temperaturesAboveZero)
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

func (c *CityTemperature) sort() []float64 {
	sort.Slice(c.temperatures, func(i, j int) bool {
		return c.temperatures[i] < c.temperatures[j]
	})
	return c.temperatures
}

// Sort
func (c CityTemperature) Len() int {
	return len(c.temperatures)
}

func (c CityTemperature) Less(i, j int) bool {
	return c.temperatures[i] < c.temperatures[j]
}

func (c CityTemperature) Swap(i, j int) {
	c.temperatures[i], c.temperatures[j] = c.temperatures[j], c.temperatures[i]
}

// Map

type NewCityTemperatures struct {
	temperatures []float64
}

func (c CityTemperature) toNewCityTemperatures() NewCityTemperatures {
	return NewCityTemperatures{c.temperatures}
}
