package main

func main() {
	// we have an array with n fields and we need tell if exist a sum of pair equal 8
	arr1 := [4]int{1, 2, 4, 9}
	target := 8

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
