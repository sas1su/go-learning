package main
import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	//go routine to work on jobs
	go func() {
		for {
			j, more := <-jobs
			//more will have false if channel closed
			if more {
				fmt.Println("recieved job:", j)
			} else {
				fmt.Println("all jobs are finished")
				done <- true
				return
			}
		}
	}()

	//sent 3 jobs
	for j := 0; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job:", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}