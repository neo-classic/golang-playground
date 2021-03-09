package main

type airTower struct {
	isRunwayFree  bool
	airplaneQueue []airplane
}

func newAirTower() *airTower {
	return &airTower{
		isRunwayFree: true,
	}
}

func (at *airTower) canLand(ap airplane) bool {
	if at.isRunwayFree {
		at.isRunwayFree = false
		return true
	}

	at.airplaneQueue = append(at.airplaneQueue, ap)
	return false
}

func (at *airTower) notifyFree() {
	if !at.isRunwayFree {
		at.isRunwayFree = true
	}
	if len(at.airplaneQueue) > 0 {
		firstAirplaneInQueue := at.airplaneQueue[0]
		at.airplaneQueue = at.airplaneQueue[1:]
		firstAirplaneInQueue.permitArrival()
	}
}
