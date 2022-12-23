package main

import (
	"fmt"
)

func evenFibNumbers() int {
	a := 1
	b := 2
	sum := 0
	for sum <= 4000000 {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, a+b
	}
	return sum
}

func main() {
	fmt.Println(evenFibNumbers())
}
