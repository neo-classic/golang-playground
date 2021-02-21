package guns

type GunInterface interface {
	setName(name string)
	setPower(power int)
	GetName() string
	GetPower() int
}
