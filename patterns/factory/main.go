package main

import (
	"fmt"
	"github.com/neo-classic/golang-playground/patterns/factory/guns"
)

func main() {
	ak47, _ := getGun(guns.Ak47)
	musket, _ := getGun(guns.Musket)

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(gun guns.GunInterface) {
	fmt.Printf("Gun: %s\n", gun.GetName())
	fmt.Printf("\tPower: %d\n\n", gun.GetPower())
}
