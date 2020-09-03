package heap

import (
	"fmt"
	"time"
)

func Heap(array []int) []int {
	start := time.Now()
	array = append([]int{0}, array...)

	for index:=2; index<len(array); index++{
		createHeap(array, index)
	}
	fmt.Println("Heap: ", array)

	for n:= len(array)-1; n>0; n-- {
		sortHeap(array, n)
	}

	elapsed := time.Since(start)
	fmt.Println("Heap sort elapsed time: ", elapsed)
	return array
}

func createHeap(array []int, n int) {
	temp := array[n]
	i := n

	for i > 1 && temp > array[i/2] {
		array[i] = array[i/2]
		i = i/2
	}
	array[i] = temp
}

func sortHeap(array []int, n int) {
	root := array[1]
	array[1] = array[n]
	i := 1

	for i*2 <= n-1  {
		j := i*2

		if j+1 >= n {
			if array[i] < array[j] {
				array[i], array[j] = array[j], array[i]
				i = j
			} else {
				i = n
			}
		}

		if array[j] >= array[j + 1] {
			if array[i] <= array[j] {
				array[i], array[j] = array[j], array[i]
				i = j
			} else {
				i = n
			}
		} else {
			if array[i] < array[j + 1] {
				array[i], array[j + 1] = array[j+1], array[i]
				i = j + 1
			} else {
				i = n
			}
		}
	}
	array[n] = root
}
