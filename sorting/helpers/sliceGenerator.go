package helpers

import (
	"math/rand"
	"time"
)

func GenerateSlice(size int) []int {
	slice := make([]int, 0, size)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice = append(slice, rand.Intn(999)-rand.Intn(999))
	}

	return slice
}
