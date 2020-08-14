package Lomuto

// the best case (highly unlikely)

import (
	"fmt"
	"time"
)

func Lomuto(array []int) []int {
	start := time.Now()

	quickSortSplit(array, 0, len(array)-1)

	elapsed := time.Since(start)
	fmt.Println("Lomuto Quicksort elapsed time: ", elapsed)
	return array
}

func quickSortSplit(array []int, low int, high int) {
	if (low < high) {
		var newPivot = partition(array, low, high)

		quickSortSplit(array, low, newPivot-1)
		quickSortSplit(array, newPivot+1, high)
	}
}

func partition(array []int, low int, high int) int {
	var pivot = high
	var i = low-1

	for j:=low; j<high; j++ {
		if (array[j] < array[pivot]) {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}

	array[high], array[i+1] = array[i+1], array[high]
	return i+1
}
