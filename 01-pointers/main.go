package main

import (
	"flag"
	"fmt"
)

// swap pointer value
func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
	fmt.Println(*a, *b)
}

// check pinter type
func pointerType() {
	str := "this is string"
	fmt.Printf("str type: %T\n", str)

	pStr := &str
	fmt.Printf("pstr type: %T\n", pStr)

	pStrValue := *pStr
	fmt.Printf("pStrValue type:%T\n", pStrValue)
	fmt.Printf("pStrValue vale:%s\n", pStrValue)
}

// check pointer address
func pointerAddr() {
	str := "this is string"
	strAddr := &str
	newStr := new(string)
	*newStr = *strAddr
	fmt.Println(*newStr)
}

// new pointer
func makePointer(str string) {
	s := new(string)
	*s = str
	fmt.Println(*s)
}

func main() {
	x, y := 1, 2
	swap(&x, &y)
	makePointer("this is pointer")
	pointerType()
	pointerAddr()

	// flag
	var mode = flag.String("mode", "", "process mode")
	flag.Parse()
	fmt.Println(*mode)
}
