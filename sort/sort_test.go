package sort

import (
	"strconv"
	"testing"
)

func TestSort(t *testing.T) {
	floatPosArr := []float64{32.1, 10.5, 36.8, 21.4, 62.1, 50.4}
	testBubbleSortFloat(t, floatPosArr)

	floatNegArr := []float64{-1, -3, -5, -7, -9, -5}
	testBubbleSortFloat(t, floatNegArr)

	floatPosNegArr := []float64{-6.8, -5.7, -4.2, 0.0, 5.6, 5.9, 7.1, 8.4, 2.6, 4.3, 1.1}
	testBubbleSortFloat(t, floatPosNegArr)

	floatSameArr := []float64{1.0, 1.0, 1.0, 1.0, 1.0, 1.0}
	testBubbleSortFloat(t, floatSameArr)

	intNegArr := []int{-1, -3, -5, -7, -9, -5}
	testMergeSortInt(t, intNegArr)
}

func BenchmarkMergeSort(b *testing.B) {
	intNegArr := []int{-1, -3, -5, -7, -9, -5}
	benchmarkMergeSort(b, intNegArr)
}

func BenchmarkBubbleSort(b *testing.B) {
	floatNegArr := []float64{-1, -3, -5, -7, -9, -5}
	benchmarkBubbleSort(b, floatNegArr)
}

func testBubbleSortFloat(t *testing.T, arr []float64) {
	BubbleSortFloat64(arr)
	for i := 0; i < (len(arr) - 1); i++ {
		if arr[i] > arr[i+1] {
			t.Error("Expected index " + strconv.Itoa(i) + " to be smaller than index " + strconv.Itoa(i+1))
		}

	}
}

func testMergeSortInt(t *testing.T, arr []int) {
	MergeSort(arr)
	for i := 0; i < (len(arr) - 1); i++ {
		if arr[i] > arr[i+1] {
			t.Error("Expected index " + strconv.Itoa(i) + " has " + strconv.Itoa(arr[i]) +
				" to be smaller than index " + strconv.Itoa(i+1) + " has " + strconv.Itoa(arr[i+1]))
		}
	}
}

func benchmarkBubbleSort(b *testing.B, arr []float64) {
	for i := 0; i < b.N; i++ {
		BubbleSortFloat64(arr)
	}
}

func benchmarkMergeSort(b *testing.B, arr []int) {
	for i := 0; i < b.N; i++ {
		MergeSort(arr)
	}
}
