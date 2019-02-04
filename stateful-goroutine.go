package main
import (
	"fmt"
	"time"
	"sync/atomic"
	"math/rand"
)

type readOp struct {
	key int
	resp chan int
}

type writeOp struct {
	key int
	val int
	resp chan bool  //to provide true resp
}

func main() {
	//record read and write ops
	var readOps uint64
	var writeOps uint64

	//channel for  read and write communication
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	//go routine owns the map repeatedly select on reads and writes channel and respond the message 
	//as they arrive
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	//statrt 100 goroutine to start to issue read ops to state go routine and 
	for i := 0; i <= 100; i++ {
		go func() {
			for {
				read := &readOp{
					key: rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	//same for write ops
	for i := 0; i < 10; i++ {
		go func() {
			for {
				write := &writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readFinalops := atomic.LoadUint64(&readOps)
	fmt.Println("readops: ", readFinalops)
	writeFinalops := atomic.LoadUint64(&writeOps)
	fmt.Println("writeops: ", writeFinalops)
	
}