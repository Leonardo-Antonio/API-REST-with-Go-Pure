package handler

import (
	"github.com/Leonardo-Antonio/api-rest-mysql/model"
)

// Storage .
type Storage interface {
	GetAll() ([]model.Alumn, error)
	GetByID(int) (model.Alumn, error)
	Create(model.Alumn) error
	Update(int, model.Alumn) error
	Delete(int) error
}
