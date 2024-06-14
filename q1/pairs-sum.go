package main

import "fmt"

func main() {
	arr1 := []int{-3, 1, 5, 4, 6, 11}
	target := 8
	existPerfectPair(arr1, target)
	fmt.Printf("The answer is: %v ", existPerfectPair(arr1, target))
}

func existPerfectPair(ar []int, target int) bool {
	left, right := 0, len(ar)-1

	for left < right {
		currentSum := ar[left] + ar[right]

		if currentSum == target {
			return true
		} else if currentSum > target {
			right++
		} else {
			left++
		}
	}
	return false
}
