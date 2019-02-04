package main
import "fmt"
import "time"

func main() {
	i := 1
	fmt.Println("write ", i , " as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its the weekend")
	default:
		fmt.Println("Its a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() > 12:
		fmt.Println("Its afternoon")
	default:
		fmt.Println("its beforenoon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am bool")
		case int:
			fmt.Println("I am int")
		default:
			fmt.Printf("Sorry I may be %T %s\n", t, i)
		}
	}
	whatAmI(1)
	whatAmI(true)
	whatAmI("hey")
}