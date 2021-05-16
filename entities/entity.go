package entities

import (
	"errors"
	"github.com/google/uuid"
)

type ID = uuid.UUID

//NewID create a new entity ID
func NewID() ID {
	return ID(uuid.New())
}

//StringToID convert a string to an entity ID
func StringToID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}


var ErrInvalidEntity = errors.New("Invalid Entity")