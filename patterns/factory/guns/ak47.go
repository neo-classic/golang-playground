package guns

const Ak47 = "ak47"

type ak47 struct {
	gun
}

func NewAk47() GunInterface {
	return &ak47{
		gun: gun{
			name:  "AK47",
			power: 4,
		},
	}
}
