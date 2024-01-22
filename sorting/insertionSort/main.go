package main

import "fmt"

// slow
func main() {
	sample := []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
	fmt.Println("Result of insertion sort: ", insertionSort(sample))
}

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			fmt.Println(arr)
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}

	return []int{}
}
