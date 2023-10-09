package arrays_slices

func Sum(arr []int) int {
	sum := 0

	for _, num := range arr {
		sum += num
	}

	return sum
}

func SumAll(arrays ...[]int) []int {

	var returnArray []int

	for _, array := range arrays {
		returnArray = append(returnArray, Sum(array))
	}

	return returnArray
}
