package main

import "fmt"

func main() {
	a := 0
	times := 3
	c := make(chan bool)

	for i := 0; i < times; i++ {
		go func() {
			a++
			c <- true
		}()
	}

	for i := 0; i < times; i++ {
		<-c
	}
	fmt.Printf("a = %d\n", a)
}
