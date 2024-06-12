package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	sq1 := square{sideLength: 2}
	tr1 := triangle{height: 2,
		base: 3}

	printArea(sq1)
	printArea(tr1)
}

func (tri triangle) getArea() float64 {
	return tri.base * tri.height * 0.5
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}
