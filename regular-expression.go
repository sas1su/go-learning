package main
import "fmt"
import "regexp"
import "bytes"

func main() {
	//match is true or false
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	//compile
	r, _ := regexp.Compile("p([a-z]+)ch")
	//check the match
	fmt.Println(r.MatchString("peach"))
	//find matched string
	fmt.Println(r.FindString("peach punch"))
	//find index of matched string start and end
	fmt.Println(r.FindStringIndex("peach punch"))
	//find the match and the submatch 
	fmt.Println(r.FindStringSubmatch("peach punch"))
	//fnd stringSubmatch Index
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	//find all substring match
	fmt.Println(r.FindAllString("peach pinch punch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", 1))
	//drop sring argument and use byte
	fmt.Println(r.Match([]byte("peach")))
	fmt.Println(r.Find([]byte("peach"))) //gives the slice of char byte
	//use mustcompile when using constant variable assignment
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)
	//replace in the regexp
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	//replace func variable allow to use a function to transfor the string
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}