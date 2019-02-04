package main
import (
	"fmt"
	"os"
)

func main() {
	//cmd with Args
	cmdWargs := os.Args
	fmt.Println(cmdWargs)
	//args without command
	argsWOcmd := os.Args[1:]
	fmt.Println(argsWOcmd)
	args := os.Args[3]
	fmt.Println(args)
}