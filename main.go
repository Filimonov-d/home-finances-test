package main

import (
	"fmt"
	"math"
	//"github.com/gin-gonic/gin"
)

type Shape interface {
	Area() float64
}

type Results interface {
	Expense() float64
	Price() float64
}

type Square struct {
	height float64
	width  float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

type Satin struct { //satingips
	price   float64 //2$ - kg
	expense float64 //1.5kg - metr
}

type Putty struct { //po derevu shpaklevka
	price   float64 //12$ - kg
	expense float64 //1kg - metr
}

type Plaster struct { //
	price   float64 //1$ - kg
	expense float64 //1kg - m2
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (s Square) Area() float64 {
	return s.height * s.width
}

func (s Satin) Expense() float64 {
	var sh Shape

	sh = Circle{radius: 5}
	return sh.Area() * s.expense

}

func (p Putty) Expense() float64 {
	var sh Shape

	sh = Circle{radius: 5}
	return sh.Area() * p.expense
}

func (d Plaster) Expense() float64 {
	var sh Shape = Circle{radius: 5}
	return sh.Area() * d.expense
}

func (d Plaster) Price() float64 {
	var sh Shape = Circle{radius: 5}
	return sh.Area() * d.price
}

func (s Satin) Price() float64 {
	var sh Shape
	sh = Circle{radius: 5}
	return sh.Area() * s.price
}

func (p Putty) Price() float64 {
	var sh Shape

	sh = Circle{radius: 5}
	return sh.Area() * p.price
}

func PrintArea(s Shape) {
	fmt.Printf("%T's area is %0.2f\n", s, s.Area())
}

func PrintResult(r Results) {
	fmt.Printf("Expense is %0.2fkg.\n", r.Expense())
	fmt.Printf("It costs - %0.2f$.\n", r.Price())
}

func main() {

	s := Square{width: 9, height: 7}
	PrintArea(s)

	c := Circle{radius: 5}
	PrintArea(c)

	sa := Satin{price: 2, expense: 1.5}
	pu := Putty{price: 12, expense: 1}
	pl := Plaster{price: 1, expense: 1}
	PrintResult(sa)
	PrintResult(pu)
	PrintResult(pl)

}
