package main

func main() {

	arr := []int{5, 2, 4, 6, 1, 3}
	SelectionSort(arr)

	for i := 0; i < len(arr); i++ {
		println(arr[i])
	}
}

func SelectionSort(arr []int) []int {
	for i := range arr {
		lower := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[lower] {
				lower = j
			}
		}
		arr[i], arr[lower] = arr[lower], arr[i]
	}

	return arr
}
