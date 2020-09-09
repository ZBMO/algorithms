package bubblesort

import (
	"fmt"
	"time"
)

func BubbleSort(numbers []int) ([]int) {
	start := time.Now()
	for i := len(numbers) - 1; i > 0; i-- {
		swapped := false
		for j := 0; j < i; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
				swapped = true
			}
		}
		if !swapped {break}
	}
	elapsed := time.Since(start)
	fmt.Println("Bubble sort elapsed time: ", elapsed)
	return numbers
}


