package main
import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	sec := now.Unix()
	p := fmt.Println
	p(sec)
	nano := now.UnixNano()
	p(nano)
	millisec := nano / 1000000
	p(millisec)

	p(time.Unix(sec, 0))
	p(time.Unix(0, nano))
}