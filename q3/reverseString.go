package main

import "fmt"

func main() {
	fmt.Println(reverse("abacaxi"))
	fmt.Println(reverse2("abacaxi----2"))
	fmt.Println(reverse3("abacaxi-----3"))
}

func reverse(str string) string {
	var result string
	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}
	return result
}

func reverse2(str string) string {
	runes := []rune(str)

	for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
		runes[l], runes[r] = runes[r], runes[l]
	}

	return string(runes)
}

func reverse3(str string) string {
	var result string

	for _, v := range str {
		result = string(v) + result
	}

	return result
}
