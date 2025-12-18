package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printDigit(n int) {
	// start at '0' and increment n times
	ch := '0'
	for i := 0; i < n; i++ {
		ch++
	}
	z01.PrintRune(ch)
}

func printNumber(n int) {
	tens := n / 10
	ones := n % 10

	printDigit(tens)
	printDigit(ones)
}

func main() {
	points := &point{}
	setPoint(points)

	printStr("x = ")
	printNumber(points.x)
	printStr(", y = ")
	printNumber(points.y)
	z01.PrintRune('\n')
}
