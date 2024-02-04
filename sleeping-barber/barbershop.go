package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientChan      chan string
	Open            bool
}

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients", barber)

		for {
			if len(shop.ClientChan) == 0 {
				color.Yellow("there is nothing to do, %s is sleeping", barber)
				isSleeping = true
			}

			client, shopOpen := <-shop.ClientChan

			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes up and starts cutting %s's hair", barber, client)
					isSleeping = false
				}

				shop.cutHair(barber, client)

			} else {
				// the shop is closed
				shop.sendBarberHome(barber)
				return
			}
		}
	}()
}

func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s is done cutting %s's hair", barber, client)
}

func (shop *BarberShop) sendBarberHome(barber string) {
	color.Cyan("%s is going home", barber)
	shop.BarbersDoneChan <- true
	shop.NumberOfBarbers--
}

func (shop *BarberShop) CloseShopForDay() {
	color.Cyan("The barbershop is closing for the day")
	shop.Open = false
	close(shop.ClientChan)

	for a := 1; a <= shop.NumberOfBarbers; a++ {
		<-shop.BarbersDoneChan
	}

	close(shop.BarbersDoneChan)
	color.Green("------------------------------------------------------")
	color.Green("All barbers have gone home, the shop is closed for the day")
}

func (shop *BarberShop) addClient(client string) {
	color.Green("***** %s arrives! *****", client)
	if shop.Open {
		select {
		case shop.ClientChan <- client:
			color.Blue("%s takes a seat in the waiting room", client)
		default:
			color.Red("Sorry %s, the waiting room is full", client)
		}
	} else {
		color.Red("Sorry %s, the barbershop is closed", client)
	}
}
