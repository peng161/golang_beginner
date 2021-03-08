package main

import (
	"github.com/peng161/golang_beginner/sort/bubblesort"
)

func main() {
	floatArr := []float64{32.1, 10.5, 36.8, 21.4, 62.1, 50.4}
	intArr := []int{32, 48, 12, 31, 25, 28}
	bubblesort.LoopThruFloat64(floatArr)
	bubblesort.LoopThruInt(intArr)
}
