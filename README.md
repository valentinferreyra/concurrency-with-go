# Concurrency Programming

Exercises learned using Go, where I practice the following concurrency concepts:

- WaitGroup
- Mutex
- Race Conditions
- Channels


**Statements of the exercises that are solved in this project:**

## The Dining Philosophers
Five philosophers live in a house together, and they always dine together at the same table, sitting in the same place.
They always eat a special king of spaghetti which requires two forks. 
There are two forks next to each plate, wich means that no two neighbours can be eating at the same time.

## The Sleeping Barber 
A barber goes to work in a barbershop with a waiting room with a fixed number of seats.
If no one is in the waiting room, the barber goes to sleep.
When a client shows up, if there are no seats available, he or she leaves.
If there is a seat available, and the barber is sleeping, the client wakes the barber up and gets a hair cut.
If the barber is busy, the client takes a seat and waits his or her turn.
Once the shop closes, no more clients are allowed in, but the barber has to stay until everyone who is waiting gets a hair cut.
