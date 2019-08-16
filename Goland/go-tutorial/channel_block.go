package main

import (
	"fmt"
	"time"
)

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func f1(in chan int) {
	fmt.Println(<-in)
}
func main() {
	// ch1 := make(chan int)
	// go pump(ch1)
	// go suck(ch1)
	//time.Sleep(1e9)
	// fmt.Println(<-ch1)
	out := make(chan int, 0)
	out <- 2
	go f1(out)
	time.Sleep(1e9)
}
