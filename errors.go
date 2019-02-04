package main
import "errors"
import "fmt"

func f1(arg int) (int, error) {
	if arg ==  42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

//define custum arg struct for explcit error
type argError struct {
	arg int
	prob string
}
//method for struct errors
func (e *argError) Error() string {
	//fmt.Sprintf("%d - %s", e.arg, e.prob)
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}
//func on Error method
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "Can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	//func f1
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed", e)
		} else {
			fmt.Println("f1 worked", r)
		}
	}
	//func f2
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed", e)
		} else {
			fmt.Println("f2 worked", r)
		}
	}
	//programatic use error with instance
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}