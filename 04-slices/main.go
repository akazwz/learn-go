package main

import "fmt"

func main() {
	letters := []string{"A", "B", "C"}
	fmt.Println(letters)
	fmt.Println(cap(letters))
	fmt.Println(len(letters))
	newLetters := append(letters, "D")
	fmt.Println(newLetters)
}
