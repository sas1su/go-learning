package main
import "fmt"
import "time"

func main() {
	//start a channel to receive requests
	requests := make(chan int, 5)
	//fill the channel for its limit
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	//start a limiter channel with 200ms
	limiter := time.Tick(200 * time.Millisecond)

	//block on limiter channel to recieve requests
	for r := range requests {
		<-limiter
		fmt.Println("request ", r, time.Now())
	}

	//start a burstylimiter channel of buffer 3
	burstylimiter := make(chan time.Time, 3)
	//fill up the channel
	for i := 1; i <=3; i++ {
		burstylimiter <- time.Now()
	}
	//go routine to send value every 200ms to burstylimiter
	go func() {
		for t := range time.Tick(2 * time.Second) {
			burstylimiter <- t
		}
	}()

	//create a burstyrequest channel with buffer 5
	burstyrequester := make(chan int, 5)
	for j := 1; j <= 5; j++ {
		burstyrequester <- j
	}
	close(burstyrequester)
	for req := range burstyrequester {
		<-burstylimiter
		fmt.Println("req ", req, time.Now())
	}

}