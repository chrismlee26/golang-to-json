package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Create structs
// Fields must be capital letters to be exported
// Exported fields will be encoded/decoded to JSON
type response1 struct {
	Page   int
	Fruits []string
}

// tag on struct field declarations to customize encoded JSON key names
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// test boolean
	boolB, _ := json.Marshal(true)
	fmt.Println(string(boolB))

	// test integer
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	// test float
	floatB, _ := json.Marshal(2.34)
	fmt.Println(string(floatB))

	// test string
	stringB, _ := json.Marshal("hello")
	fmt.Println(string(stringB))

	// test slice
	sliceD := []string{"apple", "banana", "peach"}
	sliceB, _ := json.Marshal(sliceD)
	fmt.Println(string(sliceB))

	// test map
	mapD := map[string]int{"pineapple": 5, "cherry": 9}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// JSON package encode custom data types
	// Only includes exported fields in encoded output & names = keys
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "banana", "peach"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// tag on struct field declarations
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "banana", "peach"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// JSON package decode custom data types
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["pineapple, "cherry"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"pineapple": 5, "cherry": 9}
	enc.Encode(d)

}
