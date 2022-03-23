package model

import "github.com/pborman/uuid"

type Good struct {
	ID   GoodID
	Name GoodName
}

type GoodID uuid.UUID

type GoodName string

func (s GoodName) String() string {
	return string(s)
}

func (s GoodName) IsValid() bool {
	return len(s) > 0 && len(s) < 255
}
