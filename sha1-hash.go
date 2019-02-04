package main
import (
	"fmt"
	"crypto/sha1"
	"crypto/md5"
)

func main() {
	s := "sumesh navitha"
	//new to sha1
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	//print bytestring
	fmt.Println(bs)
	fmt.Printf("%x\n", bs)
	//md5
	mh := md5.New()
	mh.Write([]byte(s))
	mbs := mh.Sum(nil)
	fmt.Println(mbs)
	fmt.Printf("%x\n", mbs)
}