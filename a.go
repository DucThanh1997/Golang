package main

import "fmt"
import "math"

// Vertex là a
type Vertex struct {
	X float64
	Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Scalee(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scalee(10)
	fmt.Println("v: ", v)
	v.Scale(10)
	fmt.Println("v: ", v)
	fmt.Println(v.Abs())


	var i interface{} = "hello"
	s, ok := i.(string)
	fmt.Println(s, ok)
}