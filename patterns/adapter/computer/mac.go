package computer

import "fmt"

type Mac struct{}

func (m *Mac) InsertIntoLightingPort() {
	fmt.Println("Lighting connector is plugged into Mac machine")
}
