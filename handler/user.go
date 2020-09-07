package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Leonardo-Antonio/api-rest-mysql/certificates/authorization"

	"github.com/Leonardo-Antonio/api-rest-mysql/model"
)

type user struct {
	storage UserStorage
}

func newUser(s UserStorage) *user {
	return &user{s}
}

func (u *user) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		reponse := NewResponse(Error, "Ha realizado mal la petición", nil)
		reponse.JSON(w, http.StatusOK)
		return
	}

	data := model.User{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		reponse := NewResponse(Error, "No se pudo obtener la iformacion del body", nil)
		reponse.JSON(w, http.StatusInternalServerError)
		return
	}

	err = u.storage.Login(data)
	if err != nil {
		reponse := NewResponse(Error, "No se encontro el usuario o no existe", nil)
		reponse.JSON(w, http.StatusBadRequest)
		return
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		response := NewResponse(Error, "No se pudo generar el token", nil)
		response.JSON(w, http.StatusInternalServerError)
		return
	}

	dataToken := map[string]string{"token": token}
	repsonse := NewResponse(Message, "Ok", dataToken)
	repsonse.JSON(w, http.StatusOK)
}

func (u *user) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := NewResponse(Error, "No realizo bien la petición", nil)
		response.JSON(w, http.StatusBadRequest)
		return
	}

	data := model.User{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := NewResponse(Error, "La estructura no es valida", nil)
		response.JSON(w, http.StatusBadRequest)
		return
	}

	err = u.storage.Create(data)
	if err != nil {
		response := NewResponse(Error, "No se pudo crear el usuario", nil)
		response.JSON(w, http.StatusInternalServerError)
		return
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		response := NewResponse(Error, "No se pudo generar el token", nil)
		response.JSON(w, http.StatusInternalServerError)
		return
	}

	dataToken := map[string]string{"token": token}
	response := NewResponse(Message, "Ok", dataToken)
	response.JSON(w, http.StatusCreated)
}

func (u *user) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := NewResponse(Error, "No realizo bien la petición", nil)
		response.JSON(w, http.StatusBadRequest)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := NewResponse(Error, "Hubo problemas con el parametro de la url", nil)
		response.JSON(w, http.StatusBadRequest)
		return
	}

	data := model.User{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := NewResponse(Error, "La estructura no es valida", nil)
		response.JSON(w, http.StatusBadRequest)
		return
	}

	err = u.storage.Update(ID, data)
	if err != nil {
		response := NewResponse(Error, "No se pudo actualizar el usuario", nil)
		response.JSON(w, http.StatusInternalServerError)
		return
	}

	response := NewResponse(Message, "Ok", nil)
	response.JSON(w, http.StatusOK)
}
