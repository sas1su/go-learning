package main
import "fmt"

func main() {
	//create a queu channel
	queue := make(chan string, 2)
	queue <- "one" 
	queue <- "two"
	close(queue)

	for r := range queue {
		fmt.Println(r)
	}
}