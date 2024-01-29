package main 

import (
	"math/rand"
	"github.com/fatih/color"
	"time"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error 
}

type PizzaOrder struct {
	pizzaNumber int 
	message string
	success bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making

	// run forever or until we receive a quit notification
	// try to make pizzas
	for {
		// try to make a pizza
		// decision: did we make a pizza or not?
	}
}

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// print out a message 
	color.Cyan("The Pizzeria is open for business.")
	color.Cyan("----------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)

	// create and run consumer

	// print out the ending message
}