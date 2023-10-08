package arrays_slices

func Sum(arr [5]int) int {
	sum := 0

	for _, num := range arr {
		sum += num
	}

	return sum
}
