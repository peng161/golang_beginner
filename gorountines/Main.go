package main

import (
	"fmt"
	"sync"
)

/*
*Reference from
*https://www.youtube.com/watch?v=YS4e4q9oBaU
 */

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go goRountineOne()
		m.Lock()
		go goRountineTwo()
	}
	wg.Wait()
}

func goRountineOne() {
	fmt.Printf("Rountine 1: %v\n", counter)
	m.RUnlock()
	wg.Done()
}

func goRountineTwo() {
	//fmt.Printf("Rountine 2: %v\n", counter)
	counter++
	m.Unlock()
	wg.Done()
}
