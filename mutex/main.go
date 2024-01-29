package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	// variable for bank balance
	var bankBalance int
	var balance sync.Mutex

	// print out starting values
	fmt.Printf("starting bank balance: $%d.00\n", bankBalance)

	// define weekly revenue
	incomes := []Income{
		{Source : "Main Job", Amount : 500},
		{Source : "Gifts", Amount : 10},
		{Source : "Part time job", Amount : 50},
		{Source : "Investments", Amount : 100},
	}

	wg.Add(len(incomes))

	// loop through 52 weeks and print out how much is made; keep a running total
	for i, income := range incomes {

		go func(i int, income Income) {
			defer wg.Done()

			for week := 1; week <= 52; week++ {
				balance.Lock()
				bankBalance += income.Amount
				fmt.Printf("Week %d: you earned $%d.00 from %s\n", week, bankBalance, income.Source)
				balance.Unlock()
			}
		}(i, income)
	}

	wg.Wait()

	// print out final values
	fmt.Printf("final bank balance: $%d.00\n", bankBalance)
}