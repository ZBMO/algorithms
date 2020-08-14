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
	lomuto "algorithms/quicksort/Lomuto"
	lomutoImproved "algorithms/quicksort/LomutoImproved"
	hoare "algorithms/quicksort/Hoare"
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

	//callBubbleSort(slice)

	//Quicksorts:
	//callLomuto(slice)
	//callLomutoImproved(slice)
	callHoare(slice)
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
