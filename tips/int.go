package main

import "fmt"

func main() {
	a := 40
	fmt.Println(a)
	// 前置补零
	b := fmt.Sprintf("%02d", a)
	fmt.Println(b)

	// 后置置零
	c := 12345
	c = c / 100 * 100
	fmt.Println(c)
}
