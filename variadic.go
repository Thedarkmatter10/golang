
// concept of variadic function:= variadic function is that function varriying number of arguments
package main

import(
	"fmt"
	"strings"
)

// (...) is known as ellipsis which is used to the type of the last parameter is preceded by an ellipsis
// It indicates that the function can be called at any number of parameters of this type. 
func joinstr(element...string)string{
	return strings.Join(element, "+")
}

func main() {

// no argument
fmt.Println(joinstr())
	
// multiple arguments
fmt.Println(joinstr("satyam", "is","backend","developer"))

	
}






