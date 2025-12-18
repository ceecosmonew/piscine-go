// package main

// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	args := os.Args[1:] // skip program name

// 	for _, word := range args {
// 		for _, char := range word {
// 			z01.PrintRune(char)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }


package main
import (
 	"os"

	"github.com/01-edu/z01"
 )

func main() {

  args := os.Args[1:]

 for _, i := range args {
	for _, g := range i{

       z01.PrintRune(g)
	}
  
z01.PrintRune('\n')
 }


    

}
