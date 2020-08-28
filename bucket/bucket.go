package radix

import (
	"fmt"
	"time"
)

func Bucket(array []int) []int {
	start := time.Now()

	bucketSort(array)

	elapsed := time.Since(start)
	fmt.Println("Bucket sort elapsed time: ", elapsed)
	return array
}


func bucketSort(array []int) {
	
	//max := getMax(array)
	//
	////make bins a slice of slices
	//bins := [max][]int
	//
	//for index, value := range array {
	//	bins[value] = append(c[i], value)
	//}
	//
	//fmt.Println(bins)
}


func getMax(array []int) int {
	x := 0
	for _, n := range array {
		if n > x {x = n}
	}
	return x
}
