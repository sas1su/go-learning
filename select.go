package main
import "fmt"
import "time"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	// two go routine starts cuncorrently
	//first stop after 1s
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	//this stops after 2s
	go func() {
		time.Sleep(15 * time.Second)
		c2 <- "two"
	}()
	//wait using select, 2 goroutine to wait for msg
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("recieved:", msg1)
		case msg2 := <- c2:
			fmt.Println("recieved:", msg2)
		}
	}

}