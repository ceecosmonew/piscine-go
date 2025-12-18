package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printHelp()
		return
	}

	insertStr := ""
	orderFlag := false
	mainStr := ""

	for i := 0; i < len(args); i++ {
		arg := args[i]

		// HELP
		if arg == "--help" || arg == "-h" {
			printHelp()
			return
		}

		// --insert=VALUE
		if len(arg) >= 9 && arg[:9] == "--insert=" {
			insertStr = arg[9:]
			continue
		}

		// -i=VALUE
		if len(arg) >= 3 && arg[:3] == "-i=" {
			insertStr = arg[3:]
			continue
		}

		// -iVALUE (everything after -i)
		if len(arg) >= 2 && arg[:2] == "-i" && arg != "-i" {
			insertStr = arg[2:]
			continue
		}

		// -i VALUE (separate)
		if arg == "-i" {
			if i+1 < len(args) {
				insertStr = args[i+1]
				i++
			}
			continue
		}

		// --order or -o
		if arg == "--order" || arg == "-o" {
			orderFlag = true
			continue
		}

		// first non-flag becomes mainStr
		if mainStr == "" {
			mainStr = arg
		}
	}

	// build result
	result := mainStr + insertStr

	// order if requested (manual bubble sort)
	if orderFlag {
		r := []rune(result)
		for a := 0; a < len(r)-1; a++ {
			for b := 0; b < len(r)-1-a; b++ {
				if r[b] > r[b+1] {
					r[b], r[b+1] = r[b+1], r[b]
				}
			}
		}
		result = string(r)
	}

	// print result + newline
	for _, c := range result {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}
