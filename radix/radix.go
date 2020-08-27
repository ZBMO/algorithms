package radix

import (
	"fmt"
	"time"
)

func Radix(array []int) []int {
	start := time.Now()

	radixSort(array)

	elapsed := time.Since(start)
	fmt.Println("Radix sort elapsed time: ", elapsed)
	return array
}


func radixSort(array []int) {
	//make bins a slice of slices
	// max := getMax(array)

	// bins := make([10]make([]int))

	// for index, value := range array {
	// 	c[value] = append(c[i], )
	// }
}


func getMax(array []int) int {
	x := 0
	for _, n := range array {
		if n > x {x = n}
	}
	return x
}
