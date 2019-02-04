package main
import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2
	fmt.Println("map:", m)

	m["k3"] = 3
	delete(m, "k2")
	fmt.Println("del:", m)
	
	v1 := m["k1"]
	fmt.Println("val:", v1)

	_, prs := m["k4"]
	fmt.Println("prs:", prs)

	decl := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("decl:", decl)
}