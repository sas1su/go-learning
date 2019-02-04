package main
import "fmt"

func sum(num ...int) int {
	total := 0
	for _, num := range num {
		total += num
	}
	return total
}

func main() {
	res := sum(1, 2)
	fmt.Println(res)

	res = sum(1, 2, 3)
	fmt.Println(res)

	s := []int{1, 2, 3, 4}
	res = sum(s...)
	fmt.Println(res)
}