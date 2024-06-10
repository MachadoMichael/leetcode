package main

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(arr, 3)
	println(arr)
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		println(nums[start], nums[end])
		start++
		end--
	}
}
