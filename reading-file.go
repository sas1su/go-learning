package main
import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"bufio"
)
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	//redfile using ioutil
	dat, err := ioutil.ReadFile("/tmp/vscode-logs/sharedprocess.log")
	check(err)
	fmt.Print(string(dat))
	//get a fd to the file and read 5 byte
	f, err := os.Open("/tmp/vscode-logs/sharedprocess.log")
	check(err)
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	//seek the file as required
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 5)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d byte @ %d: %s\n", n2, o2, string(b2))
	//io.ReadAtLeast to do the same read as above
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
	//rewind to the beginging
	_, err = f.Seek(0, 0)
	check(err)
	//bufio.NewReader()
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4) )

	//close
	f.Close()
}