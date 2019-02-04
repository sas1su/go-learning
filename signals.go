package main
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//create channel for sigs for blocking
	sigs := make(chan os.Signal, 1)
	//create a channel for freeing the block, bool channel
	done := make(chan bool, 1)
	//register the channel to recieve the signal
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	//go routine blocks for the signal
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	//wait for the signal
	//on ctr + c send TERM signal
	fmt.Println("Waiting for signal")
	<-done
	fmt.Println("Exiting")

}