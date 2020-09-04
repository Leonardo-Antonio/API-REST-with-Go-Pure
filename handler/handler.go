package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Leonardo-Antonio/api-rest-mysql/model"
)

type alumn struct {
	storage Storage
}

func newAlumn(storage Storage) *alumn {
	return &alumn{storage}
}

func (a *alumn) getByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := NewResponse(Error, "Ha realizado mal la petición", nil)
		response.JSON(w, http.StatusBadRequest)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := NewResponse(Error, "Ha ocurrido un error con el parametro id", nil)
		response.JSON(w, http.StatusInternalServerError)
		return
	}

	alumn, err := a.storage.GetByID(ID)
	if err != nil {
		response := NewResponse(Error, "Ha ocurrido un error en la bd", nil)
		response.JSON(w, http.StatusInternalServerError)
		return
	}

	response := NewResponse(Message, "Ok", &alumn)
	response.JSON(w, http.StatusInternalServerError)

}

func (a *alumn) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := NewResponse(Error, "Ha realizado mal la petición", nil)
		response.JSON(w, http.StatusBadRequest)
		return
	}

	alumns, err := a.storage.GetAll()
	if err != nil {
		response := NewResponse(Error, "Error en recuperar los datos", nil)
		response.JSON(w, http.StatusInternalServerError)
		return
	}

	response := NewResponse(Message, "OK", alumns)
	response.JSON(w, http.StatusOK)

}

func (a *alumn) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := NewResponse(Error, "Ha realizado mal la petición", nil)
		response.JSON(w, http.StatusBadRequest)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := NewResponse(Error, "No se encontro el parametro id", nil)
		response.JSON(w, http.StatusInternalServerError)
		return
	}
	alumn := model.Alumn{}
	err = json.NewDecoder(r.Body).Decode(&alumn)
	if err != nil {
		response := NewResponse(Error, "No se pudo procesar los datos del boby", nil)
		response.JSON(w, http.StatusInternalServerError)
		return
	}

	err = a.storage.Update(ID, alumn)
	if err != nil {
		response := NewResponse(Error, "Error en recuperar los datos", nil)
		response.JSON(w, http.StatusInternalServerError)
		return
	}

	response := NewResponse(Message, "OK", nil)
	response.JSON(w, http.StatusOK)

}

func (a *alumn) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := NewResponse(Error, "Ha realizado mal la petición", nil)
		response.JSON(w, http.StatusBadRequest)
		return
	}

	alumn := model.Alumn{}
	err := json.NewDecoder(r.Body).Decode(&alumn)
	if err != nil {
		response := NewResponse(Error, "No se pudo procesar los datos del boby", nil)
		response.JSON(w, http.StatusInternalServerError)
		return
	}

	err = a.storage.Create(alumn)
	if err != nil {
		response := NewResponse(Error, "Error en recuperar los datos", nil)
		response.JSON(w, http.StatusInternalServerError)
		return
	}

	response := NewResponse(Message, "OK", nil)
	response.JSON(w, http.StatusOK)

}

func (a *alumn) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := NewResponse(Error, "Ha realizado mal la petición", nil)
		response.JSON(w, http.StatusBadRequest)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := NewResponse(Error, "No se encontro el parametro id", nil)
		response.JSON(w, http.StatusInternalServerError)
		return
	}

	err = a.storage.Delete(ID)
	if err != nil {
		response := NewResponse(Error, "No se puedo eliminar", nil)
		response.JSON(w, http.StatusInternalServerError)
		return
	}

	response := NewResponse(Message, "OK", nil)
	response.JSON(w, http.StatusOK)

}
