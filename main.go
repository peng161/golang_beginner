package main

import (
	"fmt"

	"github.com/peng161/golang_beginner/sort"
)

func main() {
	floatArr := []float64{32.1, 10.5, 36.8, 21.4, 62.1, 50.4}
	intArr := []int{32, 48, 12, 31, 25, 28}
	sort.BubbleSortFloat64(floatArr)
	for i := 0; i < len(floatArr); i++ {
		fmt.Printf("Index %v contains %v\n", i, floatArr[i])
	}
	sort.BubbleSortInt(intArr)
	for i := 0; i < len(intArr); i++ {
		fmt.Printf("Index %v contains %v\n", i, intArr[i])
	}
}
