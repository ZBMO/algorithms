package Hoare

import (
	"fmt"
	"time"
)

func Hoare(array []int) []int {
	start := time.Now()

	quickSortSplit(array, 0, len(array)-1)

	elapsed := time.Since(start)
	fmt.Println("Hoare Quicksort elapsed time: ", elapsed)
	return array
}

func quickSortSplit(array []int, low int, high int) {
	if low < high {
		var partitionIndex = partition(array, low, high)

		//TODO: fix this
		//this top call gets stuck in a loop when the partition returned is same as high
		quickSortSplit(array, low, partitionIndex)
		quickSortSplit(array, partitionIndex+1, high)
	}
}

func partition(array []int, low int, high int) int {
	var pivot = array[low]
	var i, j = low, high

	for {
		for array[i] <= pivot {
			i++
		}
		for array[j] >= pivot {
			j--
		}
		if i >= j {
			fmt.Println("partition returning j: ", j)
			return j
		}
		array[i], array[j] = array[j], array[i]
	}
}
