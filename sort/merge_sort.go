package sort

//MergeSort does sorting based on integer arr
func MergeSort(intArr []int) {
	var tempArr = make([]int, len(intArr)/2+1)

	if len(intArr) < 2 {
		return
	}

	mid := len(tempArr) / 2

	MergeSort(intArr[:mid])
	MergeSort(intArr[mid:])

	if intArr[mid-1] <= intArr[mid] {
		return
	}

	copy(tempArr, intArr[:mid])

	l, r := 0, mid

	for i := 0; ; i++ {
		if tempArr[l] <= intArr[r] {
			intArr[i] = tempArr[l]
			l++

			if l == mid {
				break
			}
		} else {
			intArr[i] = intArr[r]
			r++
			if r == len(intArr) {
				copy(intArr[i+1:], tempArr[l:mid])
				break
			}
		}
	}
	return
}
