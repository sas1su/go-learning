package main
import b64 "encoding/base64"
import "fmt"

func main() {
	s := "sumesh'=~"
	//has got StdEncoding
	//with EncodeToString
	sEnc := b64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(sEnc)

	//decoding is DecodeString
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))

	//URLEncoding + EncodeToString
	// + DecodeString
	uEnc := b64.URLEncoding.EncodeToString([]byte(s))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}