
package main

import "fmt"



func ptf(a *int) {


	*a = 1200
}

func main() {

	
	var x = 101

		fmt.Printf("The value of x before function call is: %d\n", x)

	
	var pa *int = &x

	
	ptf(pa)

	fmt.Printf("The value of x after function call is: %d\n", x)

}
