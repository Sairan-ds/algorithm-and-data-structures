package main

import (
	"fmt"
	"github.com/Sairan-ds/algorithm-and-data-structures/sorting/helpers"
	"time"
)

const NumberOfElements = 20

// slow
func main() {
	array := helpers.GenerateSlice(NumberOfElements)
	fmt.Println("Unsorted array:", array)
	start := time.Now()
	fmt.Println("Result of bubble sort of array:", BubbleSort(array))
	fmt.Println("Time of sorting for ", NumberOfElements, "elements: ", time.Since(start))
}

func BubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	return array
}
