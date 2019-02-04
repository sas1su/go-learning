package main
import (
	"fmt"
	"time"
)

func main() {
	//start a new timer
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C //send a value to timer
	fmt.Println("Timer1 expired")

	timer2 := time.NewTimer(time.Second)
	//expire it if not stopped
	go func() {
		<-timer2.C //send a value to timer channel
		fmt.Println("Timer2 expired")
	}()

	stop2 := timer2.Stop()
	// if stopped 
	if stop2 {
		fmt.Println("Timer2 stopped")
	}
}