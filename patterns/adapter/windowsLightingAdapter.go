package main

import (
	"fmt"
	"github.com/neo-classic/golang-playground/patterns/adapter/computer"
)

type windowsLightingAdapter struct {
	windowMachine *computer.Windows
}

func (w *windowsLightingAdapter) InsertIntoLightingPort() {
	fmt.Println("Adapter converts Lighting signal to USB")
	w.windowMachine.InsertIntoUsbPort()
}
