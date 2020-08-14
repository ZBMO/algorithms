package Hoare

import (
	"fmt"
	"os"
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
	if (low < high) {
		var err, newPivotIndex  = partition(array, low, high)
		if err != nil {fmt.Print(err); os.Exit(1)}

		quickSortSplit(array, low, newPivotIndex)
		quickSortSplit(array, newPivotIndex+1, high)
	}
}

func partition(array []int, low int, high int) (error, int) {
	var pivot = array[low]
	var i, j = low, high

	for true {
		for array[i] < pivot {
			i++
		}
		for array[j] > pivot {
			j--
		}
		if i>=j {
			return nil, j
		}
		array[i], array[j] = array[j], array[i]
	}
	return fmt.Errorf("Error in partition function, unable to create new pivot. "), j
}
