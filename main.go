package main

import "fmt"

func main() {
	rectangle := Rectangle{3, 3}
	fmt.Println(rectangle)
	rectangle.multiplyA(3)
	fmt.Println(rectangle)
}

type Rectangle struct {
	a int
	b int
}

func (rectangle Rectangle) toSquareArea() int {
	return rectangle.a * rectangle.b
}

func (rectangle *Rectangle) multiplyA(amount int) {
	rectangle.a = rectangle.a * amount
}
