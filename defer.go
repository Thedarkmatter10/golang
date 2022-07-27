package main

import (
	"fmt"
)

func main() {
	// normally  fmt runs top to bottom
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)

	// 2. deferred statements run when the function executes its final statement
	// but before the function returns
	fmt.Println(11)
	defer fmt.Println(12)
	fmt.Println(13)

	// 3. deferred function calls are LIFO
	defer fmt.Println(21)
	defer fmt.Println(22)
	defer fmt.Println(23)

	

	a := "start"
	defer fmt.Println(a) // value decided at the time the defer is called
	a = "end"
}