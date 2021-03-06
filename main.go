package main

//ns = nanosecond - (1 billionth of 1 second) / Time to execute one machine cycle by a 1 GHz microprocessor
//µs = microsecond - (1 millionth of 1 second) / Time to execute one machine cycle by an Intel 80186 microprocessor
//ms = millisecond - (1 thousandth of 1 second)
//cs = centisecond - (1 hundredth of 1 second)

import (
	bubble "algorithms/bubblesort"
	bucket "algorithms/bucket"
	count "algorithms/count"
	insertion "algorithms/insertion"
	merge "algorithms/merge"
	hoare "algorithms/quicksort/Hoare"
	lomuto "algorithms/quicksort/Lomuto"
	lomutoImproved "algorithms/quicksort/LomutoImproved"
	radix "algorithms/radix"
	selection "algorithms/selection"
	heap "algorithms/heap"
	practice "algorithms/practice"
	"fmt"
	"math/rand"
	"time"
)

func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func main() {
	var slice = generateSlice(40)

	printResponse(practice.Practice(slice))
	//callBubbleSort(slice)
	//callSelection(slice)
	//Quicksorts:
	//callLomuto(slice)
	//callLomutoImproved(slice)
	//callHoare(slice) (don't)
	//callMerge(slice)
		//only positive numbers for count
	//callCount(slice)
	// callRadix(slice)
	// callInsertion(slice)
	// callHeap(slice)

	//callBucket(slice)
}


func printResponse(sortedList []int) {
	fmt.Println("sorted list: ", sortedList)
}

func callBubbleSort(array []int) {
	printResponse(bubble.BubbleSort(array))
}

func callLomuto(array []int) {
	printResponse(lomuto.Lomuto(array))
}
func callLomutoImproved(array []int) {
	printResponse(lomutoImproved.LomutoImproved(array))
}
func callHoare(array []int) {
	printResponse(hoare.Hoare(array))
}

func callSelection(array []int) {
	printResponse(selection.Selection(array))
}

func callMerge(array []int) {
	printResponse(merge.Merge(array))
}

func callCount(array []int) {
	printResponse(count.Count(array))
}

func callRadix(array []int) {
	printResponse(radix.Radix(array))
}

func callBucket(array []int) {
	printResponse(bucket.Bucket(array))
}

func callInsertion(array []int) {
	printResponse(insertion.Insertion(array))
}

func callHeap(array []int) {
	printResponse(heap.Heap(array))
}
