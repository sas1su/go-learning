package main
import "fmt"

func sum(a int, b int) int {
	return a + b
}

func total(a, b, c int) int {
	return a + b + c
}
func main() {
	res := sum(1, 2)
	fmt.Println("1+2:", res)

	total := total(1, 2, 3)
	fmt.Println("total:", total)
}