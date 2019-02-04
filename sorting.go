package main
import "fmt"
import "sort"

func main() {
	//sort the strings
	strs := []string{"d", "a", "b"}
	sort.Strings(strs)
	fmt.Println("strs: ", strs)
	fmt.Println("String Sorted: ", sort.StringsAreSorted(strs))

	//sort ints
	ints := []int{5, 2, 6}
	sort.Ints(ints)
	fmt.Println(ints)
	fmt.Println("sorted: ", sort.IntsAreSorted(ints))

}