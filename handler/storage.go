package handler

import (
	"github.com/Leonardo-Antonio/api-rest-mysql/model"
)

// AlumnStorage .
type AlumnStorage interface {
	GetAll() ([]model.Alumn, error)
	GetByID(int) (model.Alumn, error)
	Create(model.Alumn) error
	Update(int, model.Alumn) error
	Delete(int) error
}

// UserStorage .
type UserStorage interface {
	Create(model.User) error
	Update(int, model.User) error
	Delete(int) error
	Login(model.User) error
}
