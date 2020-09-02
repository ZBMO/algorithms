package heap

import (
	"fmt"
	"time"
)

func Heap(array []int) []int {

	start := time.Now()


	array = append([]int{0}, array...)
	fmt.Println("starting array: ", array)
	for index:=1; index<len(array); index++{
		createHeap(array, index)
	}

	fmt.Println("before heap sort: ", array)
	sortHeap(array)

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

	array[1] = array[n-1]
	i := 1
	temp := array[i]
	for i*2 < n-1  {
		j := i*2
		fmt.Printf("parent: %d, left child: %d, right child: %d\n", array[i], array[j], array[j+1])
		if array[j] >= array[j + 1] {
			if temp < array[j] {
				array[i] = array[j]
				array[j] = temp
				i = j
			}
		} else {
			if temp < array[j + 1] {
				array[i] = array[j + 1]
				array[j+1] = temp
				i = j + 1
			}
		}
	}

}
