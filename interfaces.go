package main
import (
	"fmt"
	"math"
)

// interface
type geometry interface {
	area() float64
	perim() float64
}
//struct rect
type rect struct {
	height, width float64
}

//struct circle
type circle struct {
	radius float64
}

//method for rect area
func (r rect) area() float64 {
	return r.height * r.width
}

//method for rec perim
func (r rect) perim() float64 {
	return 2 * r.height + 2 * r.width
}

//method for circle area
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

//method for circle perim
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

//named method with interface
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	//call the interface
	measure(r)
	measure(c)
}