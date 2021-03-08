package sort

import (
	"strconv"
	"testing"
)

func TestBubbleSortFloat64(t *testing.T) {
	floatPosArr := []float64{32.1, 10.5, 36.8, 21.4, 62.1, 50.4}
	testAllArrValue(t, floatPosArr)

	floatNegArr := []float64{-1, -3, -5, -7, -9, -5}
	testAllArrValue(t, floatNegArr)

	floatPosNegArr := []float64{-6.8, -5.7, -4.2, 0.0, 5.6, 5.9, 7.1, 8.4, 2.6, 4.3, 1.1}
	testAllArrValue(t, floatPosNegArr)

	floatSameArr := []float64{1.0, 1.0, 1.0, 1.0, 1.0, 1.0}
	testAllArrValue(t, floatSameArr)
}

func testAllArrValue(t *testing.T, floatArr []float64) {
	BubbleSortFloat64(floatArr)
	for i := 0; i < (len(floatArr) - 1); i++ {
		if floatArr[i] > floatArr[i+1] {
			t.Error("Expected index " + strconv.Itoa(i) + " to be smaller than index " + strconv.Itoa(i+1))
		}

	}
}
