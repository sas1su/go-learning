package main
import "fmt"
import "time"

//worker will run and feed to channel when finished
func worker(done chan bool) {
	fmt.Print("working")
	seconds := 10
	for i:= 1; i <= seconds; i++ {
		fmt.Print(".")
		time.Sleep(time.Second)
	}
	//after delay send notofication to channel
	fmt.Println("done the work")
	done <- true 
}

func main() {
	done := make(chan bool, 1)
	//go routine
	go worker(done)
	//channel blocks and wait for notification
	<-done
}