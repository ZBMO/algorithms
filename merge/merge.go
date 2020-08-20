package merge

import (
"fmt"
"time"
)

func Merge(array []int) []int {
	start := time.Now()

	a := []int{1, 3, 6, 7, 10}
	b := []int{2, 4, 6, 8, 9, 11, 25, 27, 33, 60}
	fmt.Println(mergeTwo(a, b))

	elapsed := time.Since(start)
	fmt.Println("Merge sort elapsed time: ", elapsed)
	return array
}

func mergeTwo(a []int, b []int) []int {
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