package main

import "fmt"

func main() {
	fmt.Println(reverse("abacaxi"))
}

func reverse(str string) string {
	var result string
	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}
	return result
}
