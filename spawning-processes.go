package main
import (
	"fmt"
	"os/exec"
	"io/ioutil"
)

func main() {
	//data object represent external command
	dateCmd := exec.Command("date")
	//Output() helper
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	} 
	//Output helper gives byte
	fmt.Println("> data")
	fmt.Println(string(dateOut))

	//interacting stdout, stdin
	grepCmd := exec.Command("grep", "hello")
	/*explicitly grab input grepCmd.StdinPipe
				 grab output grepCmd.StdoutPipe
				 grepCmd.Start()
	*/
	grepCmdIn, _ := grepCmd.StdinPipe()
	grepCmdOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start() 
	/* 
	grepIn.Write() 		to write 
	ioutil.ReadAll(grepout) to read
	grepIn.Close() 			to close
	grepCmd.Wait()			to wait for command to finish
	*/
	grepCmdIn.Write([]byte("hello grep\nbye bye grep hello"))
	grepCmdIn.Close()
	grepByte, _ := ioutil.ReadAll(grepCmdOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepByte))
	//need exact arguments to the command
	lsCmd := exec.Command("bash", "-c", "ls -l -a -h")
	lsCmdOut, e := lsCmd.Output()
	if e != nil {
		panic(e)
	}
	fmt.Println()
	fmt.Println("> ls -l -a -h")
	fmt.Println(string(lsCmdOut))
}