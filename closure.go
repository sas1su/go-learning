package main
import "fmt"

func intSeq() func() int {
	i := 0
	return func() int{ 
		i++
		return i
	}
}

func main() {
	newint := intSeq()

	fmt.Println(newint())
	fmt.Println(newint())
	fmt.Println(newint())

	//reint
	newInt := intSeq()
	fmt.Println(newInt())
}