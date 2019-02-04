package main
import "fmt"

//factorial
// fact(3) = 3*2*1 = 7
func fact(num int) int {
	if num == 0 {
		return 1
	}
	return num * fact(num - 1)
}

func main() {
	fmt.Println(fact(7))
}