package bubblesort

import (
	"fmt"
	"time"
)

func BubbleSort(numbers []int) ([]int) {
	start := time.Now()
	var n = len(numbers)
	for k:=0; k<n; k++ {
		sweep(numbers, k)
	}
	elapsed := time.Since(start)
	fmt.Println("Bubble sort elapsed time: ", elapsed)
	return numbers
}

func BubbleSortOptimized(numbers []int) ([]int) {
	start := time.Now()
	var n = len(numbers)
	for k:=0; k<n; k++ {
		if !sweep(numbers, k) {
			break
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Optimized bubble sort elapsed time: ", elapsed)
	return numbers
}

func sweep(numbers []int, k int) bool {
	var n = len(numbers)
	swapped := false
	for i:=0; i<n-k-1; i++ {
		if numbers[i] > numbers[i+1] {
			numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
			swapped = true
		}
	}
	return swapped
}