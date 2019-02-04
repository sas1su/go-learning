package main
import "fmt"

func main() {
	//create two channels
	messages := make(chan string)
	signals := make(chan bool)

	go func() {
		messages <- "hi"
	}()

	select {
	case msg := <-messages:
		fmt.Println("sending msg:", msg)
	default:
		fmt.Println("no message recieved")
	}

	msg := <-messages
	select {
	case messages <-msg:
		fmt.Println("recieved message:", msg)
	default:
		fmt.Println("no message recieved")
	}
	
	select {
	case msg := <-messages:
		fmt.Println("recieved message:", msg)
	case sig := <-signals:
		fmt.Println("recieved signals", sig)
	default:
		fmt.Println("no activity")
	}
}