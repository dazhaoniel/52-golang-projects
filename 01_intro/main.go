package main

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"sort"
)

// Print Hello World
func hello() {
	fmt.Println("Hello World!")
}

// Concatenate first name and last name
func concatenator() {
	firstName := "Daniel"
	lastName := "Zhao"
	fmt.Println(firstName + " " + lastName)
}

// Reverse a slice (list) of numbers
func reverseNumbers() {
	// create a list
	l := []int{}
	for range [10]int{} {
		n := rand.Intn(50)
		l = append(l, n)
	}
	// reverse it
	r := [10]int{}
	for i := range l {
		r[len(l)-1-i] = l[i]
	}
	fmt.Println("Original List: ", l)
	fmt.Println("Reversed List: ", r)
}

// Go through a list of 1 to 10 and
// print whether it's an odd or even number
func evenAndOdd() {
	n := []int{}
	for i := range [11]int{} {
		n = append(n, i)
	}

	for i := range n {
		if i%2 == 0 {
			fmt.Printf("%v is even\n", i)
		} else {
			fmt.Printf("%v is odd\n", i)
		}
	}
}

// find closest pair
func closestPair() {
	n := []int{}
	for range [10]int{} {
		n = append(n, rand.Intn(50))
	}
	sort.Ints(n)

	var a, b int
	diff := math.MaxInt32
	for i := 0; i < len(n)-1; i++ {
		if (n[i+1] - n[i]) < diff {
			diff = n[i+1] - n[i]
			a = n[i]
			b = n[i+1]
		}
	}
	fmt.Println(n)
	fmt.Println(a, b)
}

// word count
func wordCount(words []string) map[string]int {
	m := make(map[string]int)
	for _, w := range words {
		// v, ok := m[w]
		// if ok {
		// 	m[w] += 1
		// } else {
		// 	m[w] = 1
		// }
		if _, ok := m[w]; ok {
			m[w] += 1
		} else {
			m[w] = 1
		}
	}
	return m
}

func readFile() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}

func main() {
	// fmt.Println(wordCount([]string{"apple", "apple", "tea", "french"}))
	readFile()
}
