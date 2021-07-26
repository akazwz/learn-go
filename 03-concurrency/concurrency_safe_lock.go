package main

import (
	"fmt"
	"sync"
)

var (
	x int64
	w sync.WaitGroup
	lock sync.Mutex
)

func add()  {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	w.Done()
}

func main() {
	w.Add(2)
	go add()
	go add()
	w.Wait()
	fmt.Println(x)
}
