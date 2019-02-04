package main
import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	//get the binary of ls
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	//make the argument slice
	//need first arg as prog name
	args := []string{"ls", "-l", "-a", "-h"}
	//make the environments available
	//need env var also
	env := os.Environ()
	//exec the using syscall
	//exec and throw errors on failure
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}