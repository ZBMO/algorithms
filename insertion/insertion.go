package count

import (
	"fmt"
	"time"
)

// Insertion returns sorted array
func Insertion(array []int) []int {
	start := time.Now()

	insertionSort(array)

	elapsed := time.Since(start)
	fmt.Println("Insertion sort elapsed time: ", elapsed)
	return array
}

func insertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		j := i - 1
		value := array[i]

		for j > -1 && array[j] > value {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = value
	}
}
