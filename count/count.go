package count

import (
	"fmt"
	"time"
)

func Count(array []int) []int {

	start := time.Now()

	count(array)

	elapsed := time.Since(start)
	fmt.Println("Selection sort elapsed time: ", elapsed)
	return array
}

func count(array []int) {
	max := getMax(array)
	c := make([]int, max+1)

	for _, i := range array {
		c[i]++
	}

	j := 0
	for index, value := range c {
		if value > 0 {
			for x := 1; x <= value; x++ {
				array[j] = index
				j++
			}
		}
	}
}

func getMax(array []int) int {
	x := 0
	for _, n := range array {
		if n > x {x = n}
	}
	return x
}