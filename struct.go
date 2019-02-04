package main
import "fmt"

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"sumeh", 35})
	fmt.Println(person{name: "navitha", age: 34})
	fmt.Println(person{name: "fred"})

	fmt.Println(&person{"Alice", 40})

	s := person{"Ann", 50}
	fmt.Println(s.name)
	sp := &s

	sp.age = 60
	fmt.Println(sp.age)
}