package main

import "fmt"

func main() {
	fmt.Println("Hello")
	floatArr := []float64{32.1, 10.5, 36.8, 21.4, 62.1, 50.4}
	loopThru(floatArr)
}

func loopThru(floatArr []float64) {
	for i := 0; i < len(floatArr); i++ {
		pushLargerValue(floatArr, i)
	}
}

func pushLargerValue(floatArr []float64, index int) {

	for i := 0; i < len(floatArr)-1-index; i++ {
		if floatArr[i] > floatArr[i+1] {
			//fmt.Printf("Larger Value: %v, %v\n", floatArr[i], len(floatArr))
			floatArr[i], floatArr[i+1] = floatArr[i+1], floatArr[i]
		}
	}
	fmt.Printf("No %v Largest Value: %v\n", index+1, floatArr[len(floatArr)-(1+index)])
}
