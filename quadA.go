package main

import (
"fmt"
"os"
)

func main() {
if len(os.Args) != 3 {
return
}

x := atoi(os.Args[1])
y := atoi(os.Args[2])
if x <= 0 || y <= 0 {
return
}

for r := 1; r <= y; r++ {
for c := 1; c <= x; c++ {
if (r == 1 || r == y) && (c == 1 || c == x) {
fmt.Print("o")
} else if r == 1 || r == y {
fmt.Print("-")
} else if c == 1 || c == x {
fmt.Print("|")
} else {
fmt.Print(" ")
}
}
fmt.Println()
}
}

func atoi(s string) int {
n := 0
for _, r := range s {
if r < '0' || r > '9' {
return 0
}
n = n*10 + int(r-'0')
}
return n
}
