package main
import "fmt"

func main() {
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if  n := 9; n > 0 {
		fmt.Println(n, "is not negative")
	} else if n < 9 {
		fmt.Println("is more than 1 didgit")
	} else {
		fmt.Println(n, " is negative")
	}
}