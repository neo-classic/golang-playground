package guns

const Musket = "musket"

type musket struct {
	gun
}

func NewMusket() GunInterface {
	return &musket{
		gun: gun{
			name:  "Musket",
			power: 1,
		},
	}
}
