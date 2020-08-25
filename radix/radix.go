package radix

import (
	"fmt"
	"time"
)

func Radix(array []int) []int {
	start := time.Now()

	radix(array, 0, len(array)-1)

	elapsed := time.Since(start)
	fmt.Println("Radix sort elapsed time: ", elapsed)
	return array
}
