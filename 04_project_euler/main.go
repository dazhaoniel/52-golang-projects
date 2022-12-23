package main

import (
	"fmt"
	"math"
)

// https://projecteuler.net/problem=2
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

// https://projecteuler.net/problem=6
func sumSquareDifference() float64 {
	var sumOfSquare float64
	var squareOfSum float64
	for i := 1; i <= 100; i++ {
		sumOfSquare += math.Pow(float64(i), 2)
		squareOfSum += float64(i)
	}
	return math.Pow(squareOfSum, 2) - sumOfSquare
}

func main() {
	fmt.Println(sumSquareDifference())
}
