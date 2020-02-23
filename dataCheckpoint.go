package main

import (
	"fmt"
	"math"
)

/// DEFINE INTERFACE
type Shape interface {
	Area()
	Perimeter() 
}

/// DEFINE STRUCTS
type Circle struct {
	r float64
}
type Rect struct {
Width , height float64
}

//Implement the two function Area() and Perimeter to calculate a Rect and Circle area and Perimeter.
func (c Circle) Area(){
	fmt.Println("circle area :", math.Pi*c.r*c.r)
}
func (c Circle) Perimeter() {
	fmt.Println("circle perimeter :", math.Pi*c.r*2)
}
func (r Rect) Area(){
	fmt.Println("rect area :", r.Width*r.height)
}
func (r Rect) Perimeter(){
	fmt.Println("rect perimeter :", r.Width*r.height)
}


func areaInterface(s Shape)  {
	s.Area()
}
func perimeterInterface(s Shape)  {
	s.Perimeter()
}
func main(){
	var c  = Circle{2}
	var r = Rect{2,4}

	areaInterface(c)
	perimeterInterface(c)

	areaInterface(r)
	perimeterInterface(r)
}

/***********************************
Notice : Circle Area = math.Pi * r * r'
Circle Perimeter = 2 * math.Pi * r'
Rect Area = width * height'
Rect Perimeter = width * height'

******************************* */