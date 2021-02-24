package computer

import "fmt"

type Windows struct{}

func (w *Windows) InsertIntoUsbPort() {
	fmt.Println("USB connector is plugged into Windows machine")
}
