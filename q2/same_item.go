package main

import "fmt"

func main() {
	arr1 := []string{"a", "b", "c", "d"}
	arr2 := []string{"x", "k", "l", "a"}
	fmt.Printf("The answer is: %v ", sameItemInTheseArrays(arr1, arr2))
}

func sameItemInTheseArrays(arr1, arr2 []string) bool {
	dic := make(map[string]bool)

	for _, v := range arr1 {
		dic[v] = true
	}

	for _, v2 := range arr2 {
		if _, exist := dic[v2]; exist {
			return true
		}
	}
	return false
}