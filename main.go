package main

import (
	bubble "algorithms/bubblesort"
	quick "algorithms/quicksort"
	"fmt"
)

var values = []int{34, 783, 12, 53, 9, 434, 123, 435, 43, 47, 864, 764, 225, 389, 612, 420, 6543, 3, 9367, 133}

func main() {
	//callBubbleSort()
	callQuickSort()
}

func printResponse(sortedList []int) {
	fmt.Println("sorted list: ", sortedList)
}

func callBubbleSort() {
	printResponse(bubble.BubbleSort(values))
}

func callQuickSort() {
	printResponse(quick.QuickSort(values))
}
