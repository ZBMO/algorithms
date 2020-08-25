package main

//notes:
//time complexity - worst, best
//in place / out of place?
//when to use it? when not to use it?
//what languages use it and where?
//

//ns = nanosecond - (1 billionth of 1 second) / Time to execute one machine cycle by a 1 GHz microprocessor
//µs = microsecond - (1 millionth of 1 second) / Time to execute one machine cycle by an Intel 80186 microprocessor
//ms = millisecond - (1 thousandth of 1 second)
//cs = centisecond - (1 hundredth of 1 second)

import (
	bubble "algorithms/bubblesort"
	"algorithms/count"
	lomuto "algorithms/quicksort/Lomuto"
	lomutoImproved "algorithms/quicksort/LomutoImproved"
	hoare "algorithms/quicksort/Hoare"
	selection "algorithms/selection"
	merge "algorithms/merge"
	radix "algorithms/radix"
	"fmt"
	"math/rand"
	"time"
)

func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(1)
	}
	return slice
}

func main() {
	var slice = generateSlice(40)

	//callBubbleSort(slice)

	//callSelection(slice)

//Quicksorts:
	//callLomuto(slice)
	//callLomutoImproved(slice)
	//callHoare(slice) (don't)

	//callMerge(slice)

	//callCount(slice)

	callRadix(slice)
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
