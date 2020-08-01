package quicksort

import (
	"fmt"
	"time"
)

func QuickSort(numbers []int) []int {
	start := time.Now()

	partition(numbers, 0, len(numbers)-1)


	elapsed := time.Since(start)
	fmt.Println("Quick sort elapsed time: ", elapsed)
	return numbers
}


func partition(array []int, low int, high int) int {
	var pivot = array[high]

	var i = -1
	for j:=low; j<=high; j++ {
		if (array[j]<=pivot) {
			i++
			swap(array, i, j)
		}

	}
	swap(array, i+1, high)
	fmt.Println("once partition array: ", array)
	fmt.Println("pivoted on: ", i+1)
	return i+1
}

func swap(array []int, low int, high int) {
	var temp = array[low]
	array[low] = array[high]
	array[high] = temp
}