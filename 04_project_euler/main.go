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

func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// https://projecteuler.net/problem=7
func tenThoughsandthPrime() int {
	c := 1
	n := 3
	for {
		if isPrime(n) {
			c += 1
		}
		if c == 10001 {
			break
		}
		n += 1
	}
	return n
}

func main() {
	fmt.Println(tenThoughsandthPrime())
}
