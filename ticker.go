package main
import "fmt"
import "time"

func main() {
	//Start a new ticker ticks every 500ms
	ticker1 := time.NewTicker(250 * time.Millisecond)

	//func to recieve the ticker channel
	go func() {
		for t := range ticker1.C {
			fmt.Println("Tick at ", t)
		}
	}()

	//stop ticker same after 1600, should stop after 3 ticks
	time.Sleep(1600 * time.Millisecond)
	ticker1.Stop()
	fmt.Println("Ticker Stopped")
}