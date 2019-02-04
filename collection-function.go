package main
import (
	"strings"
	"fmt"
)

func Index(vs []string, t string) int {
	for i, s := range vs {
		if s == t {
			return i
		}
	}
	return -1
}

func Includes(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

//any of the fruit starts with prefix
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}
//All string in slice has prefix then true
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}
//func return a slice with matched condition from predcate
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
// create slice, make the upper, assign to the slice
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
func main() {
	var strs = []string{"peach", "banana", "orange", "apple", "gree-ngrape"}
	//return index of target string t if matches 
	fmt.Println(Index(strs, "orange"))

	//return true if there is a match, call func from another
	fmt.Println(Includes(strs, "apples"))

	//check the slice for condition, fruit starts with p
	fmt.Println(Any(strs, func(vs string) bool {
		return strings.HasPrefix(vs, "p")
	}))

	//All return true if slice has all starts with "p"
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	//fil ter string starts with b and return slice
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.HasPrefix(v, "b")
	}))
	//map the upper of string in slice and return slice
	//strings.ToUpper
	fmt.Println(Map(strs, strings.ToUpper))
	
}