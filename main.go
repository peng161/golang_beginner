package main

import (
	//"github.com/peng161/golang_beginner/gorountines" //Use for go-rountines example
	"github.com/peng161/golang_beginner/sort" //Use for sort example
)

func main() {
	floatArr := []float64{32.1, 10.5, 36.8, 21.4, 62.1, 50.4}
	intArr := []int{32, 48, 12, 31, 25, 28}
	sort.BubbleSortFloat64(floatArr)
	/*
		for i := 0; i < len(floatArr); i++ {
			fmt.Printf("Index %v contains %v\n", i, floatArr[i])
		}
	*/
	sort.BubbleSortInt(intArr)
	/*
		for i := 0; i < len(intArr); i++ {
			fmt.Printf("Index %v contains %v\n", i, intArr[i])
		}
	*/

	/*
		for i := 0; i < 10; i++ {
			gorountines.WaitGroup.Add(2)
			gorountines.Mutex.RLock()
			go gorountines.GoRountineOne()
			gorountines.Mutex.Lock()
			go gorountines.GoRountineTwo()
		}
		gorountines.WaitGroup.Wait()
	*/
}
