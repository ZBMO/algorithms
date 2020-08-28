package radix

import (
	"fmt"
	"time"
	"math"
)

func Radix(array []int) []int {
	start := time.Now()

	radixSort(array)

	elapsed := time.Since(start)
	fmt.Println("Radix sort elapsed time: ", elapsed)
	return array
}


func radixSort(array []int) {
	max := getMax(array)
	bins := [10][]int{}
	numberOfPasses := recursionCountDigits(max)

	for pass:=1; pass<=numberOfPasses; pass++ {

		divisor := int(math.Pow(10, float64(pass-1)))
		for _, value := range array {
			indexOfBin := (value/divisor)%10
			bins[indexOfBin] = append(bins[indexOfBin], value)
		}

		arrayIndex := 0
		for i:=0; i<len(bins); i++ {
			bin := bins[i]
			if len(bins[i]) != 0 {
				for _, number := range bin {
					array[arrayIndex] = number
					bin = bin[1:]
					bins[i] = bin
					arrayIndex++
				}
			}
		}
		fmt.Printf("array after %d sorts: %d\n", pass, array)
	}
}


func getMax(array []int) int {
	x := 0
	for _, n := range array {
		if n > x {x = n}
	}
	return x
}

func recursionCountDigits(max int) int {
	if max < 10 {
		return 1
	} else {
		return 1 + recursionCountDigits(max/10)
	}
}
