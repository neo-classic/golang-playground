package main

type address struct {
	city   string
	street string
	zip    int
}

func NewAddress(city, street string, zip int) *address {
	return &address{
		city:   city,
		street: street,
		zip:    zip,
	}
}

func (a *address) Zip() int {
	return a.zip
}

func (a *address) SetZip(zip int) *address {
	return &address{
		city:   a.City(),
		street: a.Street(),
		zip:    zip,
	}
}

func (a *address) Street() string {
	return a.street
}

func (a *address) SetStreet(street string) *address {
	return &address{
		city:   a.City(),
		street: street,
		zip:    a.Zip(),
	}
}

func (a *address) City() string {
	return a.city
}

func (a *address) SetCity(city string) *address {
	return &address{
		city:   city,
		street: a.Street(),
		zip:    a.Zip(),
	}
}
