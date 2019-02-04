package main
import "fmt"
import "time"

func main() {
	p := fmt.Println
	now := time.Now()
	p(now)
	then := time.Date(
		2009, 01, 17,  05, 24, 10, 0, time.UTC)
	p(then)
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(now.Location())
	p(then.Weekday())

	p(then.After(now))
	p(then.Before(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)
	p(diff.Hours())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
	
}