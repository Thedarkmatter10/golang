package main

import "fmt"
       "time"

func tutree(str string){
	for w :=0; w<6; w++{
		time.sleep(1* time.second)
		fmt.Println(str)
	} 
}
func main(){
	 go tutree("hello")

	tutree("backend developer")
}