package sort_algo

// SelectionSort algorithm
func SelectionSort(data []int) Data {
	iteration := 0
	n := len(data)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if data[j] < data[minIdx] {
				minIdx = j
			}
		}
		data[i], data[minIdx] = data[minIdx], data[i]
		iteration++
	}
	return Data{
		Iteration: iteration,
		Sorted:    data,
	}
}
