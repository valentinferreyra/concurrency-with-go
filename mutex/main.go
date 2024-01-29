package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()

	m.Lock()
	msg = s
	m.Unlock()
}

func main() {
	msg = "Hello, Go!"

	var mu sync.Mutex

	wg.Add(2)
	go updateMessage("Hello, Java!", &mu)
	go updateMessage("Hello, Scala!", &mu)
	wg.Wait()

	fmt.Println(msg)
}