package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("received order #%d, making pizza...\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg, success := "", false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("making pizza #%d, it will take %d seconds\n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("pizza order #%d failed", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("pizza order #%d is late", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("pizza order #%d is ready!", pizzaNumber)
		}

		return &PizzaOrder{pizzaNumber, msg, success}
	}

	return &PizzaOrder{pizzaNumber: pizzaNumber}
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	var i = 0

	// run forever or until we receive a quit notification
	// try to make pizzas
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			// we tried to make a pizza (we sent something to the data chan)
			case pizzaMaker.data <- *currentPizza:

			// we want to quit, so send pizzaMaker.quit chan to the quitChan (a chan error)
			case quitChan := <-pizzaMaker.quit:
				// close channels
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
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
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("order #%d is out for delivery", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("the customer is really mad!")
			}
		} else {
			color.Cyan("done making pizzas...")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("error closing channel!: %v", err)
			}
		}
	}

	color.Cyan("---------------------")
	color.Cyan("done for the day...")
	color.Cyan("we made %d, but failed to make %d, with %d attempts in total.", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 9:
		color.Red("we need to improve our service, it was an awful day!")
	case pizzasFailed > 5:
		color.Yellow("it was not a very good day...")
	case pizzasFailed > 2:
		color.Yellow("it was a good day!")
	default:
		color.Green("it was a great day!")
	}
}
