package LomutoImproved

// the best case (highly unlikely)

import (
	"fmt"
	"time"
)

func LomutoImproved(array []int) []int {
	start := time.Now()

	quickSortSplit(array, 0, len(array)-1)

	elapsed := time.Since(start)
	fmt.Println("Lomuto Quicksort with median-of-three pivot elapsed time: ", elapsed)
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
	var pivot = medianOfThree(array, low, high)
	var i = low-1

	for j:=low; j<high; j++ {
		if (array[j] < pivot) {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}

	array[high], array[i+1] = array[i+1], array[high]
	return i+1
}

func medianOfThree(array []int, low int, high int) int {
	var mid = low + (high - low)/2
	if (array[mid] < array[low]) {array[mid], array[low] = array[low], array[mid]}
	if (array[high] < array[low]) {array[high], array[low] = array[low], array[high]}
	if (array[mid] < array[high]) {array[mid], array[high] = array[high], array[mid]}
	return array[high]
}
