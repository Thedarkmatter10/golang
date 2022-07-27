package main

import "fmt"

type animal struct {
	name string

	numberoflegs int
}

func main() {
	var a animal
	fmt.Println(a)

	a1 := animal{"lion", 4}
	fmt.Println("animal1: ", a1)

	a2 := animal{"human being", 2}
	fmt.Println("animal: ", a2)

}


