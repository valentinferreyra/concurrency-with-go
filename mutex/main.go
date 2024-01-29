package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	word := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	wg.Add(len(word))

	for i, w := range word {
		go printSomething(fmt.Sprintf("%d: %s", i, w), &wg)
	}

	wg.Wait()

	wg.Add(1)
	printSomething("This is the second thing I'm printing", &wg)
}