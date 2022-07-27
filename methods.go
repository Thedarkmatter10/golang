package main

import "fmt"

type bank struct {
	name   string
	branch string

	amount int
}

func (a bank) show() {

	fmt.Println("Account holder Name: ", a.name)
	fmt.Println("Branch Name: ", a.branch)

	fmt.Println("amount: ", a.amount)
}

func main() {

	d := bank{
		name:   "tanmay",
		branch: "noida",

		amount: 34000,
	}

	d.show()
}
