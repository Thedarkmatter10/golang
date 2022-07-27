package main

import "fmt"

func fibo(a int) {
	if a == 0 {
		return 0
	}

	if a == 1 {
		return 1
	}
	return fibo(a-1) + (a - 2)
}

func main() {
	var a int
	for a = 0; a < 10; a++ {
		fmt.Printf("%d ", fibo(a))
	}
}
