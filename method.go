package main
import "fmt"

type rect struct {
	width, height int
}
func (r *rect)area() int {
	return r.width * r.height
}

func (r rect) perm() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 3, height: 2}

	fmt.Println("area:", r.area() )
	fmt.Println("Perim:", r.perm() )

	rp := &r

	fmt.Println("parea:", rp.area())
	fmt.Println("pperm:", rp.perm())
}