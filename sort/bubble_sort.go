package sort

import "fmt"

func LoopThruFloat64(floatArr []float64) {
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

func LoopThruInt(intArr []int) {
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
