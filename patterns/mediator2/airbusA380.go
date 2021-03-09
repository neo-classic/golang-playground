package main

import "fmt"

type airbusA380 struct {
	meidator mediator
}

func (a *airbusA380) requestArrival() {
	if a.meidator.canLand(a) {
		fmt.Println("AirbisA380: Landing")
	} else {
		fmt.Println("AirbusA380: Waiting")
	}
}

func (a *airbusA380) departure() {
	fmt.Println("AirbusA380: Leaving")
	a.meidator.notifyFree()
}

func (a *airbusA380) permitArrival() {
	fmt.Println("AirbusA380: Arrival Permitted. Landing")
}
