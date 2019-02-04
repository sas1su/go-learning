package main
import "fmt"

func main() {
	s := []int{1, 2, 3, 4}

	//sum the value in the slice
	total := 0
	for _, num := range s {
		total += num
	}
	fmt.Println("slice num:", total)

	//fing the index of 3
	for i, v := range s {
		if v == 3 {
			fmt.Println("Index:", i)
		}
	}

	//range map, print key => val
	kv := map[string]int{"k1": 1, "k2": 2, "k3": 3}
	for k, v := range kv {
		fmt.Printf("%s -> %d\n", k, v)
	}

	//print key
	for k := range kv {
		fmt.Println("Keys:", k)
	}

	for i, c := range "sumesh" {
		fmt.Println(i, c)
	}
}