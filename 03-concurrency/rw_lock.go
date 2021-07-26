package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	a int64
	rwLock sync.RWMutex
	swg sync.WaitGroup
)

func write()  {
	rwLock.Lock()
	a = a + 1
	time.Sleep(10 * time.Millisecond)
	rwLock.Unlock()
	swg.Done()
}

func read()  {
	rwLock.RLock()
	time.Sleep(time.Millisecond)
	rwLock.RUnlock()
	swg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		swg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		swg.Add(1)
		go read()
	}

	swg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
