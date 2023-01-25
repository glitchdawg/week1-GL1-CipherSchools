package main

import "fmt"

func main() {
	//rec(4)
	rec2(4)
}

// func rec(num int) {
// 	if num == 0 {
// 		return
// 	}
// 	if num%2 == 0 {
// 		rec(num - 1)
// 		fmt.Println(num - 1)
// 	} else {
// 		rec(num - 1)
// 		fmt.Println(num - 1)
// 	}
// 	fmt.Println(num - 1)

// }
func rec2(num int) {
	if num <= 0 {
		return
	}
	rec2(num - 1)
	rec2(num - 2)
	fmt.Println(num - 1)
}
