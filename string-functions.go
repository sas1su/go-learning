package main
import s "strings"
import "fmt"

var p = fmt.Println

func main() {
	p("Contains:", s.Contains("Sumesh", "um"))
	p("Counts:", s.Count("text", "t"))
	p("Hasprefix:", s.HasPrefix("peach", "p"))
	p("HasSuffix:", s.HasSuffix("Banana", "na"))
	p("Join:", s.Join([]string{"a", "b"}, "-"))
	p("Index:", s.Index("test", "e"))
	p("Replace:", s.Replace("foo", "o", "0", -1))
	p("Replace 1:", s.Replace("foo", "o", "0", 1))
	p("Repeat:", s.Repeat("sumesh", 5))
	p("Split:", s.Split("a-b-c", "-"))
	p("Tolower:", s.ToLower("TOLOWER"))
	p("ToUpper:", s.ToUpper("ToUpper"))

	p("len:", len("sumesh"))
	p("leng char", "navitha"[1])
}