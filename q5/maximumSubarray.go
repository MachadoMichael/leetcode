package main

func main() {
	arr := []int{2, 3, -4, 5, 7, 2 - 1}
	maxSubArray(arr)
}

func maxSubArray(nums []int) int {
	max := nums[0]
	curSum := 0
	for _, n := range nums {
		if curSum < 0 {
			curSum = 0
		}

		curSum += n
		max = biggest(max, curSum)
	}

	return max

}

func biggest(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
