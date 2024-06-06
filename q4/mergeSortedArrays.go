package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 14}
	arr2 := []int{5, 6, 17, 18}
	fmt.Printf("%v ", mergeSortedArrays(arr1, arr2))
}

func mergeSortedArrays(arr1, arr2 []int) []int {
	var result []int
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}

	}

	for i < len(arr1) {
		result = append(result, arr1[i])
		i++
	}

	for j < len(arr2) {
		result = append(result, arr2[j])
		j++
	}

	return result
}
