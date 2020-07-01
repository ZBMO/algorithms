package quicksort

import (
	"fmt"
	"time"
)

func QuickSort(numbers []int) []int {
	start := time.Now()

	elapsed := time.Since(start)
	fmt.Println("Quick sort elapsed time: ", elapsed)
	return numbers
}