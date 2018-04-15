package main

import (
	"log"
	"sync"
)

func main() {
	a := 0
	times := 10000
	c := make(chan bool)

	var m sync.Mutex

	for i := 0; i < times; i++ {
		go func() {
			m.Lock()
			a++
			m.Unlock()
			c <- true
		}()
	}

	for i := 0; i < times; i++ {
		<-c
	}
	log.Printf("a = %d", a)
}
