package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isPrime(n int) bool {
	if n == 0 || n == 1 {
		return false
	}
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func checkNumber() {
	fmt.Println("Please enter a number or q to quit")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// check to see if the user wants to quit
	if strings.EqualFold(scanner.Text(), "q") {
		os.Exit(0)
	}

	// try to convert what the user typed into an int
	num, _ := strconv.Atoi(scanner.Text())

	msg := isPrime(num)
	if msg {
		fmt.Printf("%d is a prime", num)
	} else {
		fmt.Printf("%d is not a prime", num)
	}

}

func main() {
	checkNumber()
}
