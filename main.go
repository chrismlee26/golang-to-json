package main

import (
	"encoding/json"
	"fmt"
)

// Create structs
type response1 struct {
	Page   int
	Fruits []string
}

// Fields must be capital letters to be exported
// Ecported fields will be encoded/decoded to JSON
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
}
