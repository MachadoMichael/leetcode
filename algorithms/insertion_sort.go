package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	arr := []int{5, 2, 4, 6, 1, 3}
	arr = insertionSort(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func insertionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] <= arr[i+1] {
			continue
		}

		arr[i+1], arr[i] = arr[i], arr[i+1]
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}
