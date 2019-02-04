package main
import "fmt"

func val() (int, int) {
	return 3, 7
}

func main() {
	k, v := val()
	fmt.Println(k, v)

	_, v = val()
	fmt.Println(v)
}