package main
import "fmt"
        

func main(){

 age := -22

 error := fmt.Errorf("%d is negative\n age can't be negative", age)

 if age<0{
 fmt.Println(error)
 } else{
	fmt.Println("age : %d", age);
  }

}

