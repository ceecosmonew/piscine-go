package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	// Need at least 4 args: program, "-c", number, file...
	if len(args) < 4 || args[1] != "-c" {
		return
	}

	// Parse positive integer after -c (manual parse to avoid strconv)
	count := 0
	for _, r := range args[2] {
		if r < '0' || r > '9' {
			return
		}
		count = count*10 + int(r-'0')
	}

	files := args[3:]
	errorOccurred := false
	anyPrinted := false // true if we've printed either an error or some file output

	for _, filename := range files {
		content, err := os.ReadFile(filename)
		if err != nil {
			// Print error line alone (no header)
			fmt.Printf("%s\n", err)
			errorOccurred = true
			anyPrinted = true
			continue
		}

		// If something was printed before (error or success), print a blank line
		if anyPrinted {
			fmt.Printf("\n")
		}

		// Print header only when multiple files are provided
		if len(files) > 1 {
			fmt.Printf("==> %s <==\n", filename)
		}

		// Compute the start index for the tail
		start := len(content) - count
		if start < 0 {
			start = 0
		}

		// Print last bytes (note: content includes newline if file ends with it)
		fmt.Printf("%s", string(content[start:]))

		anyPrinted = true
	}

	if errorOccurred {
		os.Exit(1)
	}
}
