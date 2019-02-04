package main
import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	//create func with channel only accepts
	ping(pings, "passed message")

	// create func with channel receive and send
	pong(pings, pongs)
	fmt.Println(<-pongs)
}