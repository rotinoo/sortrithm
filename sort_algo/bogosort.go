package sort_algo

import (
	"math/rand"
)

func isSorted(data []int) bool {
	for i := 1; i < len(data); i++ {
		if data[i-1] > data[i] {
			return false
		}
	}
	return true
}

func shuffle(data []int) {
	for i := range data {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

func randomslice(length int) []int {
	array := make([]int, length)
	for i := range array {
		array[i] = rand.Intn(10)
	}
	return array
}

type Data1 struct {
	Data      []int `json:"array"`
	Iteration int   `json:"iteration"`
	Sorted    []int `json:"sorted"`
}

type Data struct {
	Iteration int   `json:"iteration"`
	Sorted    []int `json:"sorted"`
}

func BogoSort(data []int) Data {

	iteration := 0
	for !isSorted(data) {
		shuffle(data)
		iteration++
	}
	return Data{
		Iteration: iteration,
		Sorted:    data,
	}
}

func BogoSortRand(length int) Data1 {
	ArBefore := randomslice(length)
	ArAfter := make([]int, len(ArBefore))
	copy(ArAfter, ArBefore)
	iteration := 0
	for !isSorted(ArAfter) {
		shuffle(ArAfter)
		iteration++
	}
	return Data1{
		Data:      ArBefore,
		Iteration: iteration,
		Sorted:    ArAfter,
	}
}
