package main

import "github.com/neo-classic/golang-playground/patterns/adapter/computer"

func main() {
	client := &client{}
	mac := &computer.Mac{}

	client.insertLightingConnectorIntoComputer(mac)

	windowsMachine := &computer.Windows{}
	windowsMachineAdapter := &windowsLightingAdapter{
		windowMachine: windowsMachine,
	}
	client.insertLightingConnectorIntoComputer(windowsMachineAdapter)
}
