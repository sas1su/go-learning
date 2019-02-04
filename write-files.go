package main
import (
	"fmt"
	"bufio"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//write some byte to the file
	b1 := []byte("hello\nthere")
	err := ioutil.WriteFile("/tmp/dat1", b1, 0644)
	check(err)

	//os.Create for more granular file operation
	f, err := os.Create("/tmp/dat2")
	check(err)
	defer f.Close()

	//using the ponter to write
	b2 := []byte{110, 120, 114, 113, 121}
	n1, err := f.Write(b2)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, b2)

	//using WriteString
	n3, err := f.WriteString("Hello you there\n")
	check(err)
	fmt.Printf("%d bytes written\n", n3)
	f.Sync()
	//using bufio.NewWriter for Buffered write
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("Something else\n")
	check(err)
	fmt.Println("written bytes", n4)
	w.Flush()


}