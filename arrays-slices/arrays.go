package arrays_slices

import "slices"

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

		sum := 0

		for _, elem := range array {
			sum += elem
		}

		slices.Insert(returnArray, sum)
	}

	return returnArray
}
