package main

type airplane interface {
	requestArrival()
	departure()
	permitArrival()
}
