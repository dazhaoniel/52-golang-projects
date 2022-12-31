package main

import (
	"fmt"
	"sync"
)

func printNames(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	names := []string{
		"Daniel",
		"John",
		"Javier",
		"Jennifer",
		"Andrea",
		"Ada",
		"Grace",
		"Jillian",
		"Michael",
	}

	wg.Add(len(names))

	for i, v := range names {
		go printNames(fmt.Sprintf("%d. %s", i, v), &wg)
	}

	wg.Wait()
}
