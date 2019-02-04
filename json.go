package main
import "encoding/json"
import "fmt"
import "os"

type response1 struct {
	Page int
	Fruits []string
}
//struct with tag
type response2 struct {
	Page int        `json:"page"`
	Fruits []string `json:"fruits"`
}
func main() {
	//encode basic datatypes
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	strB, _ := json.Marshal("peach")
	fmt.Println(string(strB))
	slcD := []string{"peach", "apple", "banana"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	mapD := map[string]int{"peach": 5, "apple": 2, "banana": 10}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	//custom datatype encode with json
	res1D := &response1{
		Page: 1, 
		Fruits: []string{"peach", "apple", "banana"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
//use the tags for the data type
	res2D := &response2{
		Page: 1,
		Fruits: []string{"peach", "apple", "banana"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))
//decoding json from go data type
	byt := []byte(`{"num":6.3,"strs":["a","b","c"]}`)
	//provide a datatype for json decoder to put the data
	//key string and variable val 
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	fmt.Println(dat["num"])
	//casting of nested data
	strs := dat["strs"].([]interface{})
	fmt.Println(strs[1])
	//using the custom type to encode
	str := `{"Page":1, "Fruits": ["a", "b", "c"]}`
	res := &response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	//without byte and string. Streamed to os.Stdout
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "peach": 7}
	enc.Encode(d)

}