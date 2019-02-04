package main
import (
	"fmt"
	"flag"
)

func main() {
	//flag.String
	wordPtr := flag.String("word", "sumesh", "Its a string") 
	//flag.Int
	intPtr := flag.Int("numb", 42, "Its a number")
	//flag.Bool
	boolPtr := flag.Bool("v", true, "Its a bool")
	//passing var defines using ptr
	var svar string
	flag.StringVar(&svar, "svar", "bar", "Its a var")

	flag.Parse()

	if *boolPtr {
		fmt.Println("word:", *wordPtr)
		fmt.Println("Number:", *intPtr)
		fmt.Println("Bool:", *boolPtr)
		fmt.Println("svar:", svar)
		//positional arguments, slice of args
		fmt.Println("Pos Arg:", flag.Args()[0])
	}
}