package main
import "fmt"

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)

	fmt.Printf("%T\n", p)
	fmt.Printf("%t\n", true)

	fmt.Printf("%d\n", 123)

	fmt.Printf("%b\n", 32)
	fmt.Printf("%c\n", 101)

	fmt.Printf("%x\n", 450)

	fmt.Printf("%f\n", 48.6)
}