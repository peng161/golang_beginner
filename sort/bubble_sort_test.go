package sort

import "testing"

func TestBubbleSortFloat64(t *testing.T) {
	floatArr := []float64{32.1, 10.5, 36.8, 21.4, 62.1, 50.4}
	BubbleSortFloat64(floatArr)
	if floatArr[len(floatArr)-1] != 62.1 {
		t.Error("Expected Biggest Value to be 62.1")
	} else if floatArr[0] != 10.5 {
		t.Error("Expected Smallest Value to be 10.5")
	}
}
