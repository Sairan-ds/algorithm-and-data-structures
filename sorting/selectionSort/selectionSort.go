package main

import (
	"fmt"
	"github.com/Sairan-ds/algorithm-and-data-structures/sorting/helpers"
)

// slow
func main() {
	testArray := helpers.GenerateSlice(30)

	fmt.Println("Array for selection sort", testArray)
	fmt.Println("Result of selection sort", selectionSort(testArray))
}

func selectionSort(arr []int) (resultArray []int) {
	for range arr {
		smallest := findSmallest(arr)
		resultArray = append(resultArray, arr[smallest])
		arr[smallest] = arr[len(arr)-1]
		arr = arr[:len(arr)-1]
	}

	return resultArray
}

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0

	for i, v := range arr {
		if v < smallest {
			smallest = v
			smallestIndex = i
		}
	}

	return smallestIndex
}
