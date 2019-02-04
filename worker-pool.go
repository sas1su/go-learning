package main
import "fmt"
import "time"

func worker(id int, jobs <-chan int, results chan<- int) {
	//recieve the job on range on channel jobs
	for j := range jobs {
		fmt.Println("worker ", id, "started job ", j)
		time.Sleep(time.Second)
		fmt.Println("worked ", id, "finished job", j)
		results <- j * 2 //send the results to results channel
	}
}
func main() {
	//two channel for jobs and results
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//worked go routine with jobs to start the jobs and result to get the results
	//will start 3 worker
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	//send five jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	//get the results from the channel
	for r := 1; r <= 5; r++ {
		<-results
	}

}