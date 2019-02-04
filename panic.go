package main
import "os"

func main() {
	panic("Somethig wrong")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}