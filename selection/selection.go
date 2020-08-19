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
	for h, _ := range array {
		i:= h
		for j:=h; j<len(array); j++ {
			if array[j] < array[i] {
				i=j
			}
		}
		if h != i {
			array[h], array[i] = array[i], array[h]
		}
	}
}

func selectionSortCocktail(array []int) {
	//TODO: fix this
	for h, _ := range array {
		var swapped = false
		i:= h
		m:= h
		max:=len(array)-1
		for j:=h; j<max; j++ {
			if array[j] < array[i] {
				i = j
			}
			if (array[j] > array[m]) {
				m = j
			}
		}
		if h != i {
			swapped = true
			array[h], array[i] = array[i], array[h]
		}
		if m != max {
			swapped = true
			array[m], array[max] = array[max], array[m]
		}
		max--
		if swapped == false{
			return
		}
	}
}

