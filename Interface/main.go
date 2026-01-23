package main

import "fmt"

type Shape interface {
	Area() float64
}
type Measure interface{
	Perimeter() float64
}
type Geometry interface{
	Shape
	Measure           	
}

type Rectangle struct {
	len, wid float64
}
func (r Rectangle) Area() float64 {
	return r.len * r.wid
}
func(r Rectangle) Perimeter() float64{
	return r.len + r.wid
}
func Desrcibe(g Geometry) {
	fmt.Println("The Area is :", g.Area())
	fmt.Println("The Perimeter is :", g.Perimeter())
}

func main() {
	rec := Rectangle{
		len: 12,
		wid: 10,
	}
	Desrcibe(rec)
}
