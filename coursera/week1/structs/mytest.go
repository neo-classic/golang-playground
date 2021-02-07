package main

import "fmt"

type Person struct {
	Id      int
	Name    string
	Address Address
}

type Address struct {
	City     string
	Street   string
	Address1 string
}

func (a Address) GetFullAddress() string {
	return fmt.Sprintf("%s, %s, %s", a.City, a.Street, a.Address1)
}

func (a *Address) SetFullAddress(city, street, address1 string) {
	a.City = city
	a.Street = street
	a.Address1 = address1
}

type SomeMap map[string]string

func (sm SomeMap) Add(key, val string) {
	sm[key] = val
}

func (sm SomeMap) Print() {
	fmt.Printf("%#v\n", sm)
}

func (sm SomeMap) Count() int {
	return len(sm)
}

func main() {
	person := Person{
		Id:   1,
		Name: "Yuriy B",
		Address: Address{
			City:     "Volgograd",
			Street:   "Test st",
			Address1: "123",
		},
	}

	fmt.Printf("%#v\n", person)
	fmt.Println(person.Address.GetFullAddress())
	person.Address.SetFullAddress("Saratov", "Street st", "2232")
	fmt.Println(person.Address.GetFullAddress())

	fmt.Println("==========================================")

	sm := SomeMap{"name": "Yuriy", "address": "Addr"}
	sm.Print()
	sm.Add("www", "eee")
	sm.Print()
	fmt.Println(sm.Count())
}
