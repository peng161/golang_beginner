package main

import "fmt"

func main() {

	floatArr := []float64{32.1, 10.5, 36.8, 21.4, 62.1, 50.4}
	intArr := []int{32, 48, 12, 31, 25, 28}
	fmt.Println("Float Arr")
	loopThruFloat64(floatArr)
	fmt.Println("Int Arr")
	loopThruInt(intArr)
}

func loopThruFloat64(floatArr []float64) {
	for i := 0; i < len(floatArr); i++ {
		pushLargerValueFloat64(floatArr, i)
	}
}

func pushLargerValueFloat64(floatArr []float64, index int) {

	for i := 0; i < len(floatArr)-1-index; i++ {
		if floatArr[i] > floatArr[i+1] {
			//fmt.Printf("Larger Value: %v, %v\n", floatArr[i], len(floatArr))
			floatArr[i], floatArr[i+1] = floatArr[i+1], floatArr[i]
		}
	}
	fmt.Printf("No %v Largest Value: %v\n", index+1, floatArr[len(floatArr)-(1+index)])
}

func loopThruInt(intArr []int) {
	for i := 0; i < len(intArr); i++ {
		pushLargerValue(intArr, i)
	}
}

func pushLargerValue(intArr []int, index int) {

	for i := 0; i < len(intArr)-1-index; i++ {
		if intArr[i] > intArr[i+1] {
			//fmt.Printf("Larger Value: %v, %v\n", floatArr[i], len(floatArr))
			intArr[i], intArr[i+1] = intArr[i+1], intArr[i]
		}
	}
	fmt.Printf("No %v Largest Value: %v\n", index+1, intArr[len(intArr)-(1+index)])
}
