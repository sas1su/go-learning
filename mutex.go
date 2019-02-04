package main
import (
	"fmt"
	"time"
	"sync"
	"sync/atomic"
	"math/rand"
)

func main() {
	//create state map for read and write ops
	state := make(map[int]int)
	//keep positive uint64 var to keep the ops
	var readOps uint64
	var writeOps uint64
	//mutex to sync the access to state
	mutex := &sync.Mutex{}
	
	//start 100 go routine to read from state
	//lock the mutex to exclusive access to state
	//update read every millisecond
	for i := 1; i <= 100; i++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)  //generate a key
				//lock the mutex, read value and incerement counter 1
				mutex.Lock()
				total += state[key]
				mutex.Unlock()  //unlock the mutex
				atomic.AddUint64(&readOps, 1)
				//wait millisecond
				time.Sleep(time.Millisecond)
			}
		}()
	}

	//same way lock mutex and exclusive write to state
	for i := 0; i < 10; i++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	
	//wait some time
	time.Sleep(time.Second)
	//read the finalval of ops
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readops: ", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeops: ", writeOpsFinal)

	//lock the mutex and get the state
	mutex.Lock()
	fmt.Println("state: ", state)
	mutex.Unlock()
}