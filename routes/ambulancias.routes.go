package routes

import (
	"encoding/json"

	"errors"
	"net/http"

	"gorm.io/gorm"

	"github.com/Francisco9djul/go-apirest/db"
	"github.com/Francisco9djul/go-apirest/models"
	"github.com/gorilla/mux"
)

func GetAmbulanciasHandler(w http.ResponseWriter, r *http.Request) { //GET
	var ambulancias []models.Ambulancia

	// Cargar todos las ambulancias junto con sus accidentes
	if err := db.DB.Preload("Accidentes").Find(&ambulancias).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error al obtener las ambulancias"))
		return
	}

	json.NewEncoder(w).Encode(&ambulancias)
}

func GetAmbulanciaHandler(w http.ResponseWriter, r *http.Request) { //GET {id}
	var ambulancia models.Ambulancia
	params := mux.Vars(r)

	if err := db.DB.Preload("Accidentes").First(&ambulancia, params["id"]).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Ambulancia no encontrado"))
		return
	}

	json.NewEncoder(w).Encode(&ambulancia)
}

func PostAmbulanciaHandler(w http.ResponseWriter, r *http.Request) { //POST
	var ambulancia models.Ambulancia

	json.NewDecoder(r.Body).Decode(&ambulancia)

	// Verificar si el chofer relacionado existe
	var choferExistente models.Chofer
	if err := db.DB.First(&choferExistente, ambulancia.ChoferID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusBadRequest) // 400
			w.Write([]byte("Chofer no encontrado"))
			return
		}
	}

	// Verificar si el paramédico relacionado existe
	var paramedicoExistente models.Paramedico
	if err := db.DB.First(&paramedicoExistente, ambulancia.ParamedicoID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusBadRequest) // 400
			w.Write([]byte("Paramédico no encontrado"))
			return
		}
	}

	createdAmbulancia := db.DB.Create(&ambulancia)
	err := createdAmbulancia.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&ambulancia)
}

func PutAmbulanciaHandler(w http.ResponseWriter, r *http.Request) {
	var ambulancia models.Ambulancia
	params := mux.Vars(r)
	db.DB.First(&ambulancia, params["id"])

	if ambulancia.ID == 0 {

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("ambulancia no encontrado"))
		return
	}

	var updatedAmbulancia models.Ambulancia
	json.NewDecoder(r.Body).Decode(&updatedAmbulancia)

	// Verificar si el chofer relacionado existe
	var choferExistente models.Chofer
	if err := db.DB.First(&choferExistente, updatedAmbulancia.ChoferID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusBadRequest) // 400
			w.Write([]byte("Chofer no encontrado"))
			return
		}
	}

	// Verificar si el paramédico relacionado existe
	var paramedicoExistente models.Paramedico
	if err := db.DB.First(&paramedicoExistente, updatedAmbulancia.ParamedicoID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusBadRequest) // 400
			w.Write([]byte("Paramédico no encontrado"))
			return
		}
	}

	ambulancia.Patente = updatedAmbulancia.Patente
	ambulancia.Inventario = updatedAmbulancia.Inventario
	ambulancia.Vtv = updatedAmbulancia.Vtv
	ambulancia.Seguro = updatedAmbulancia.Seguro
	ambulancia.Estado = updatedAmbulancia.Estado
	ambulancia.Estado = updatedAmbulancia.Estado
	ambulancia.ChoferID = updatedAmbulancia.ChoferID
	ambulancia.ParamedicoID = updatedAmbulancia.ParamedicoID

	db.DB.Save(&ambulancia)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ambulancia)
}

func DeleteAmbulanciaHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var ambulancia models.Ambulancia
	// Buscar el ambulancia en la base de datos
	if err := db.DB.First(&ambulancia, params["id"]).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("ambulancia no encontrado"))
		return
	}

	ambulancia.Estado = "baja"

	if err := db.DB.Save(&ambulancia).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error al actualizar el ambulancia"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ambulancia eliminado lógicamente"))
}
