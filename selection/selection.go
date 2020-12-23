package selection

import (
	"fmt"
	"time"
)

func Selection(array []int) []int {
	start := time.Now()

	selectionSort(array)

	elapsed := time.Since(start)
	fmt.Println("Selection sort elapsed time: ", elapsed)
	return array
}

func selectionSort(array []int) {
	for i, _ := range array {
		 min := i
		for j :=i+1; j<len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		if min != i {
			array[min], array[i] = array[i], array[min]
		}
	}
}
