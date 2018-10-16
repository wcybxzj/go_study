package main

import "fmt"
import "math"

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

//对类型进行扩展
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}

	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("r1", r1.area())
	fmt.Println("r2", r2.area())
	fmt.Println("c1", c1.area())
	fmt.Println("c2", c2.area())
}
