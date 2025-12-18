package main

import (
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		return
	}

	aStr := args[0]
	op := args[1]
	bStr := args[2]

	a, okA := atoi(aStr)
	b, okB := atoi(bStr)

	if !okA || !okB {
		return
	}

	switch op {
	case "+":
		res, ok := add(a, b)
		if ok {
			printInt(res)
		}
	case "-":
		res, ok := sub(a, b)
		if ok {
			printInt(res)
		}
	case "*":
		res, ok := mul(a, b)
		if ok {
			printInt(res)
		}
	case "/":
		if b == 0 {
			printStr("No division by 0")
			return
		}
		res := a / b
		printInt(res)
	case "%":
		if b == 0 {
			printStr("No modulo by 0")
			return
		}
		res := a % b
		printInt(res)
	default:
		return
	}
}

// ---------------------
// Helper functions
// ---------------------

func atoi(s string) (int64, bool) {
	sign := int64(1)
	i := 0
	if len(s) == 0 {
		return 0, false
	}
	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}

	var num int64 = 0
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		num = num*10 + int64(s[i]-'0')
	}

	return num * sign, true
}

func add(a, b int64) (int64, bool) {
	r := a + b
	if (b > 0 && r < a) || (b < 0 && r > a) {
		return 0, false
	}
	return r, true
}

func sub(a, b int64) (int64, bool) {
	r := a - b
	if (b < 0 && r < a) || (b > 0 && r > a) {
		return 0, false
	}
	return r, true
}

func mul(a, b int64) (int64, bool) {
	r := a * b
	if a != 0 && r/a != b {
		return 0, false
	}
	return r, true
}

func printInt(n int64) {
	// Handle zero quickly
	if n == 0 {
		os.Stdout.Write([]byte("0\n"))
		return
	}

	// Handle negative sign without newline
	if n < 0 {
		os.Stdout.Write([]byte{'-'})
		n = -n
	}

	var digits [20]byte
	i := 0
	for n > 0 {
		digits[i] = byte(n%10) + '0'
		n /= 10
		i++
	}

	for j := i - 1; j >= 0; j-- {
		os.Stdout.Write([]byte{digits[j]})
	}
	os.Stdout.Write([]byte{'\n'})
}

func printStr(s string) {
	os.Stdout.Write([]byte(s))
	os.Stdout.Write([]byte{'\n'})
}
