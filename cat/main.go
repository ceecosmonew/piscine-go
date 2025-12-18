package main

import (
	"bufio"
	"io"
	"os"
)

func printString(s string) {
	for _, r := range s {
		os.Stdout.Write([]byte(string(r)))
	}
}

func printFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(os.Stdout, f)
	return err
}

func printStdin() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		printString(line)
		if err == io.EOF {
			break
		}
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printStdin()
		return
	}

	for _, file := range args {
		err := printFile(file)
		if err != nil {
			// Print EXACT error format without fmt
			msg := "ERROR: " + err.Error() + "\n"
			os.Stderr.Write([]byte(msg))
			os.Exit(1)
		}
	}
}
