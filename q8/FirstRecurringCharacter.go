package main

import (
	"fmt"
)

func main() {
	var arr []int = []int{1, 1, 2, 3, 4, 5}
	res, err := FirstRecurringCharacter(arr)
	if err == true {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}

func FirstRecurringCharacter(arr []int) (int, bool) {
	dic := make(map[int]bool)

	for _, v := range arr {
		if dic[v] != true {
			dic[v] = true
		} else {
			return v, false
		}
	}

	return 0, true

}
