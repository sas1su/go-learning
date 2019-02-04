package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	//bufio.NewScanner to scan unbuffered stdin
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(scanner)
	//go through each line sep by \n
	for scanner.Scan() {
		tuc := strings.ToUpper(scanner.Text())
		fmt.Println(tuc)
	}

	//capture the error from scanner.Err()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}