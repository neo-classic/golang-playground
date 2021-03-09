// https://levelup.gitconnected.com/the-mediator-pattern-in-go-344ee5c8c2f4

package main

func main() {
	airTower := newAirTower()
	boeing737 := &boeing737{
		mediator: airTower,
	}
	airbusA380 := &airbusA380{
		meidator: airTower,
	}

	boeing737.requestArrival()
	airbusA380.requestArrival()
	boeing737.departure()
	airbusA380.departure()
}
