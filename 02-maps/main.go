package main

import "fmt"

// make map
func makeMap() {
	dict := make(map[string]string)

	dict["a"] = "apple"
	dict["b"] = "banana"

	value, ok := dict["a"]
	if ok {
		fmt.Println("value a existed:", value)
	} else {
		fmt.Println("value a not existed")
	}

	valueC, ok := dict["c"]
	if ok {
		fmt.Println("value c existed:", valueC)
	} else {
		fmt.Println("value c not existed")
	}
}

// range map and delete map by key
func optionMap()  {
	dict := make(map[string]string)
	dict["a"] = "apple"
	dict["b"] = "banana"
	dict["c"] = "cherry"
	delete(dict, "d")
	for key, value := range dict {
		fmt.Println(key, value)
	}
}

func main() {
	makeMap()
	optionMap()
}
