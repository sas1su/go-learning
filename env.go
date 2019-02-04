package main
import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("Foo", "1")
	fmt.Println("Foo:", os.Getenv("Foo"))
	fmt.Println("Bar:", os.Getenv("Bar"))

	fmt.Println()
	for _, e := range os.Environ() {
		pairs := strings.Split(e, "=")
		fmt.Println(pairs[0])
	}
}