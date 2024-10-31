package main

// func main() {
//
// 	arr := []int{5, 2, 4, 6, 1, 3, 9, 7}
// 	mergeSort(arr)
// 	for n := range arr {
// 		println(n)
// 	}
// }
//
// func mergeSort(arr []int) []int {
// 	if len(arr) <= 1 {
// 		return arr
// 	}
//
// 	size := len(arr)
// 	var splittedSlice [][]int
// 	splittedSlice = append(splittedSlice, arr)
//
// 	for len(splittedSlice) < size {
// 		for index, array := range splittedSlice {
// 			if len(array) <= 1 {
// 				continue
// 			}
//
// 			splittedSlice = append(splittedSlice[:index], splittedSlice[index+1:]...)
// 			splittedSlice = append(splittedSlice, array[:len(array)/2])
// 			splittedSlice = append(splittedSlice, array[len(array)/2:])
// 		}
// 	}
//
// 	for len(splittedSlice) > 1 {
// 		for i := 0; i < len(splittedSlice); i++ {
// 			splittedSlice[i] = merge(splittedSlice[i], splittedSlice[i+1])
// 			splittedSlice = append(splittedSlice[:i+1], splittedSlice[i+2:]...)
// 		}
// 	}
//
// 	arr = splittedSlice[0]
// 	return arr
// }
//
// func merge(a, b []int) []int {
// 	var arr []int
// 	for i := 0; i < len(a); i++ {
// 		if a[i] < b[i] {
// 			arr = append(arr, a[i])
// 			arr = append(arr, b[i])
// 		} else {
// 			arr = append(arr, b[i])
// 			arr = append(arr, a[i])
// 		}
// 	}
// 	return arr
//
// }

// need use a second array to sort the array
func merge(left, right []int) []int {
	var result []int
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func printSlice(arr []int) {
	for _, n := range arr {
		println(n)
	}

	println("finish")
}

func main() {
	arr := []int{500, 25, 4, 6, 1, 332, 9, 7, 0, 8}
	sorted := mergeSort(arr)
	printSlice(sorted)

}
