package main

import (
    "github.com/01-edu/z01"

   "piscine"
)

func main() {
	a := piscine.SplitWhiteSpaces("Hello how are you?")
	PrintWordsTables(a)
}


func PrintWordsTables(a []string) string{

	for _, sett := range  a {

		for _, sec := range sett {
			
			z01.PrintRune(sec)
		} 
		z01.PrintRune('\n')
	}
    return 
}