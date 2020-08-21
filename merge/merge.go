package merge

import (
"fmt"
"time"
)

func Merge(array []int) []int {
	start := time.Now()

	newArray := mergeSort(array)

	elapsed := time.Since(start)
	fmt.Println("Merge sort elapsed time: ", elapsed)
	return newArray
}

func mergeSort(array []int) []int {
	if len(array) == 1 {
		return array
	}
	mid := len(array)/2
	return twoWayMerge(mergeSort(array[:mid]), mergeSort(array[mid:]))
}

func twoWayMerge(a []int, b []int) []int {
	var c []int
	i, j:= 0, 0
	for i<len(a) && j<len(b) {
		if a[i] <= b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}
	for i<len(a) {
		c = append(c, a[i])
		i++
	}
	for j<len(b)  {
		c = append(c, b[j])
		j++
	}
	return c
}