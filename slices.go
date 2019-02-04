package main
import "fmt"

func main() {
	s := make([]string, 4)
	fmt.Println("emp:", s)
	s[0] = "ab"
	s[1] = "a"
	s[2] = "b"
	s[3] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	//slice
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[3:]
	fmt.Println("sl3:", l)

	t := []int{1, 2, 3, 4}
	fmt.Println("dcl:", t)

	// towd slice, variale inner length
	towD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerlength := i + 1
		towD[i] = make([]int, innerlength)
		for j := 0; j < innerlength; j++ {
			towD[i][j] = i + j
		}
	}
	fmt.Println("towD:", towD)
}