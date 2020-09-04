package bubblesort

import (
	"fmt"
	"time"
)

func BubbleSort(numbers []int) ([]int) {
	start := time.Now()
	var n = len(numbers)
	for j:=n-1; j>0; j-- {
		if !sweep(numbers, j) {
			break
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Bubble sort elapsed time: ", elapsed)
	return numbers
}

func sweep(numbers []int, j int) bool {
	swapped := false
	for k:=0; k<j; k++ {
		if numbers[k] > numbers[k+1] {
			numbers[k], numbers[k+1] = numbers[k+1], numbers[k]
			swapped = true
		}
	}
	return swapped
}
