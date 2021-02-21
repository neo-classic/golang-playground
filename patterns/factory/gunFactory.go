package main

import (
	"fmt"
	"github.com/neo-classic/golang-playground/patterns/factory/guns"
)

func getGun(gunName string) (guns.GunInterface, error) {
	switch gunName {
	case guns.Ak47:
		return guns.NewAk47(), nil
	case guns.Musket:
		return guns.NewMusket(), nil
	default:
		return nil, fmt.Errorf("Wrong gunName passed")
	}
}
