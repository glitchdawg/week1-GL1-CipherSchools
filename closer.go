package main

import "fmt"

func main() {
	// number := 10
	var number int
	number = 10
	fmt.Println(number)

	//store a function as a value to variable
	var getInt func(x int) int
	getInt = func(x int) int {
		fmt.Println("In a 1st function")
		return 20 + x
	}
	g(getInt, 8)
	// getInt = func(x int) int {
	// 	fmt.Println("In a 2nd function")
	// 	return 10 + x
	// }
	var y func()
	y = func() {
		package main

		import "fmt"
		
		func main() {
			// number := 10
			var number int
			number = 10
			fmt.Println(number)
		
			//store a function as a value to variable
			var getInt func(x int) int
			getInt = func(x int) int {
				fmt.Println("In a 1st function")
				return 20 + x
			}
			g(getInt, 8)
			// getInt = func(x int) int {
			// 	fmt.Println("In a 2nd function")
			// 	return 10 + x
			// }
			var y func()
			y = func() {
				fmt.Print(9)
	}
	y()
	getInt = func(x int) int {
		fmt.Println("In a 1st function")
		return 20 + x
	}
	j := getInt(19)
	fmt.Println(j)
}
func g(getInt func(int) int, u int) {
	getInt(78)
}

//function is a first class member in golang
