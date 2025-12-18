package main

 import (
 	"fmt"
 	
)


func main()  {
	s := []string{"Hello", "How" ,"Are" , "12you"}
	if len(s) == 0 {
		fmt.Println("false")
	}


	for _, me := range s {
		

	 	if me[0] >= 'A' && me[0] <= 'Z' {
			fmt.Println("true")
		}

		if rune(me[0]) >= '0' {
			fmt.Println("true")
		}
		if me[0] >= 'a' && me[0] <= 'z'{
			fmt.Println("false")
		}

	}


	
}