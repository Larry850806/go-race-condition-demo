package main

import (
	"log"
)

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
	log.Printf("a = %d", a)
}
