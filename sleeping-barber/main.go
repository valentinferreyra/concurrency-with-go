package main

import (
	"fmt"

	"github.com/fatih/color"

	"math/rand"
	"time"
)

var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	// seed our random number generator
	rand.Seed(time.Now().UnixNano())

	// print welcome message
	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("------------------------------------")

	// create channels if we need any
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	// create the barbershop
	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientChan:      clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	color.Green("The barbershop is open for business!")

	// add barbers to the barbershop
	shop.addBarber("Frank the Barber")
	shop.addBarber("Gerrard the Busy Barber")
	shop.addBarber("Gill the Slow Barber")
	shop.addBarber("Mario the Fast Barber")
	shop.addBarber("Shawn the Bad Barber")

	// start the barbershop as a goroutine
	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.CloseShopForDay()
		closed <- true
	}()

	// add clients to the barbershop
	i := 1

	go func() {
		for {
			// get a random nomber with average arrival rate
			randomMillisecond := rand.Intn(2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Duration(randomMillisecond) * time.Millisecond):
				shop.addClient(fmt.Sprintf("Client %d", i))
				i++
			}
		}
	}()

	// block until the barbershop is closed
	<-closed
	time.Sleep(5 * time.Second)
}
