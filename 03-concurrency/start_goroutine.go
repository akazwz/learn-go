package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func hello(i int)  {
	defer wg.Done()
	fmt.Println("Hello, Goroutine!", i)
}

func printA()  {
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func printB()  {
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func startGoroutine()  {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
}

func goMaxProgress()  {
	runtime.GOMAXPROCS(2)
	go printA()
	go printB()
	time.Sleep(time.Second)
}