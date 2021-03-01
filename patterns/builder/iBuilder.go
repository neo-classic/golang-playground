// Builder interface
package main

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

const (
	typeNormal = "normal"
	typeIgloo  = "igloo"
)

func NewBuilder(builderType string) iBuilder {
	if builderType == typeNormal {
		return &normalBuilder{}
	}

	if builderType == typeIgloo {
		return &iglooBuilder{}
	}

	return nil
}
