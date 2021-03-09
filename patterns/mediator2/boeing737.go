package main

import "fmt"

type boeing737 struct {
	mediator mediator
}

func (a *boeing737) requestArrival() {
	if a.mediator.canLand(a) {
		fmt.Println("Boeing737: Landing")
	} else {
		fmt.Println("Boeing737: Waiting")
	}
}

func (a *boeing737) departure() {
	fmt.Println("Boeing737: Leaving")
	a.mediator.notifyFree()
}

func (a *boeing737) permitArrival() {
	fmt.Println("Boeing737: Arrival Permitted. Landing")
}
