package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var a int32
	times := 10000
	c := make(chan bool)

	for i := 0; i < times; i++ {
		go func() {
			atomic.AddInt32(&a, 1)
			c <- true
		}()
	}

	for i := 0; i < times; i++ {
		<-c
	}
	fmt.Printf("a = %d\n", a)
}
