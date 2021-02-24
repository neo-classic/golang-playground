package main

import (
	"github.com/neo-classic/golang-playground/patterns/adapter/computer"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
)

type client struct{}

func (c *client) insertLightingConnectorIntoComputer(com computer.ComputerInterface) {
	fmt.Println("Client inserts Lighting connector into computer")
	com.InsertIntoLightingPort()
}
