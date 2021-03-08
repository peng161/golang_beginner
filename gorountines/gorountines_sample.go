package gorountines

import (
	"fmt"
	"sync"
)

/*
*Reference from
*https://www.youtube.com/watch?v=YS4e4q9oBaU
 */

//WaitGroup to ensure only one group write at a time
var WaitGroup = sync.WaitGroup{}

//Mutex to sync read write
var Mutex = sync.RWMutex{}

var counter = 0

//GoRountineOne print the counter
func GoRountineOne() {
	fmt.Printf("Rountine 1: %v\n", counter)
	Mutex.RUnlock()
	WaitGroup.Done()
}

//GoRountineTwo add the counter
func GoRountineTwo() {
	//fmt.Printf("Rountine 2: %v\n", counter)
	counter++
	Mutex.Unlock()
	WaitGroup.Done()
}
