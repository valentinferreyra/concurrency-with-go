package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

// list of all philosophers
var philosophers = []Philosopher{
	{name: "Plato", rightFork: 0, leftFork: 4},
	{name: "Socrates", rightFork: 1, leftFork: 0},
	{name: "Aristotle", rightFork: 2, leftFork: 1},
	{name: "Pascal", rightFork: 3, leftFork: 2},
	{name: "Locke", rightFork: 4, leftFork: 3},
}

// define variables
var hunger = 3 // how main times a philosopher will eat
var eatTime = 1 * time.Second
var thingTime = 3 * time.Second
var sleepTime = 1 * time.Second

var orderMutex sync.Mutex
var orderFinished []string

func main() {
	// print out a welcome message
	fmt.Println("welcome to the dining philosophers' problem")
	fmt.Println("------------------------------------------")
	fmt.Println("the table is empty.")

	time.Sleep(sleepTime)

	// start the meal
	dine()

	// print out a goodbye message
	fmt.Println("the table is empty.")

	time.Sleep(sleepTime)
	fmt.Printf("the philosophers ate in the following order: %v\n", orderFinished)
}

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// map of forks
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// start the meal
	for i := 0; i < len(philosophers); i++ {
		// fire off a goroutine for the current philosopher
		go dinningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()

}

func dinningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	// seat the philosopher at the table
	fmt.Printf("%s is seated at the table.\n", philosopher.name)
	seated.Done()

	seated.Wait()

	// eat three times
	for i := hunger; i > 0; i-- {
		// get a lock on both forks
		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s has the right fork.\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s has the left fork.\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s has the left fork.\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s has the right fork.\n", philosopher.name)
		}

		fmt.Printf("\t%s has both forks is eating.\n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking.\n", philosopher.name)
		time.Sleep(thingTime)

		// release the locks
		forks[philosopher.rightFork].Unlock()
		forks[philosopher.leftFork].Unlock()

		fmt.Printf("\t%s put down the forks.\n", philosopher.name)
	}

	fmt.Printf("%s is satisfied and is leaving the table.\n", philosopher.name)

	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher.name)
	orderMutex.Unlock()
}
