package main
import "fmt"
import "time"

func main() {
	//create two channel
	c1 := make(chan string, 1)
	//go rotine to send message to c1
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	//wait and select on c1
	select {
	case r := <-c1:
		fmt.Println(r)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout 1")
	}

	//create channel c2
	c2 := make(chan string, 1)
	//go routine to send message c2
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	//wait and select on c2
	select {
	case r := <-c2:
		fmt.Println(r)
	case <-time.After(4 * time.Second):
		fmt.Println("timeout 2")
	}
}