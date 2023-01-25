package main

import "fmt"

func fib(number int) int {
	if number == 0 || number == 1 {
		return 1

	}
	result := fib(number-1) + fib(number-2)
	return result

}

func fact(numbers int) int {
	if numbers == 1 || numbers == 0 {
		return 1
	}
	ans := numbers * fact(numbers-1)
	return ans
}
func main() {
	fmt.Println(fib(4))
	fmt.Println(fact(5))
}
