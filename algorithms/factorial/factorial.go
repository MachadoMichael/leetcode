package main

func main() {
	println(recursiveFactorial(5))
	println(iterativeFactorial(5))
}

func recursiveFactorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * recursiveFactorial(num-1)
}

func iterativeFactorial(num int) int {
	for i := num - 1; i > 0; i-- {
		num *= i
	}
	return num
}
