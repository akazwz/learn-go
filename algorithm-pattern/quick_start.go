package main

import "fmt"

func main() {
	index := indexOf("this is stack", "ck")
	fmt.Println(index)
}

// strStr 28
func indexOf(haystack, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(needle) > len(haystack) {
		return 0
	}

	var i, j int

	for i = 0; i < len(haystack)-len(needle)+1; i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if len(needle) == j {
			return i
		}
	}
	return -1
}
