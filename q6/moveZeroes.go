package main

import "fmt"

func main() {
	arr := []int{0, 1, 0, 3, 12}

	fmt.Println(moveZeroes(arr))
}

func moveZeroes(nums []int) []int {
	lastNonZeroIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNonZeroIndex], nums[i] = nums[i], nums[lastNonZeroIndex]
			lastNonZeroIndex++

		}
	}
	return nums
}
