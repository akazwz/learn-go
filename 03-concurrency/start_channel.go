package main

import "fmt"

func receive(c chan int) {
	res := <-c
	fmt.Printf("receive %v success\n", res)
}

// no buffer sync
func startChan() {
	ch := make(chan int)
	go receive(ch)
	ch <- 50
	fmt.Println("send success")
}

func startBufferedChan() {
	ch := make(chan int, 1) //cap 1 buffered chan
	ch <- 10
	fmt.Println("send success")
}

func rangeChan() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()

	for i := range ch2 {
		fmt.Println(i)
	}
}

func main() {
	//startChan()
	//startBufferedChan()
	rangeChan()

}
