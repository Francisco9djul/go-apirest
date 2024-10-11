package routes

import (
	"encoding/json"

	"net/http"

	"github.com/Francisco9djul/go-apirest/db"
	"github.com/Francisco9djul/go-apirest/models"
	"github.com/gorilla/mux"
)

func GetParamedicosHandler(w http.ResponseWriter, r *http.Request) { //GET
	var paramedicos []models.Paramedico

	// Cargar todos los Paramedicos junto con sus ambulancias
	if err := db.DB.Preload("Ambulancias").Find(&paramedicos).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error al obtener los paramedicos"))
		return
	}

	json.NewEncoder(w).Encode(&paramedicos)
}

func GetParamedicoHandler(w http.ResponseWriter, r *http.Request) { //GET {id}
	var paramedico models.Paramedico
	params := mux.Vars(r)

	if err := db.DB.Preload("Ambulancias").First(&paramedico, params["id"]).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Paramedico no encontrado"))
		return
	}

	json.NewEncoder(w).Encode(&paramedico)
}

func PostParamedicoHandler(w http.ResponseWriter, r *http.Request) { //POST
	var paramedico models.Paramedico

	json.NewDecoder(r.Body).Decode(&paramedico)

	createdParamedico := db.DB.Create(&paramedico)
	err := createdParamedico.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&paramedico)
}

func PutParamedicoHandler(w http.ResponseWriter, r *http.Request) {
	var paramedico models.Paramedico
	params := mux.Vars(r)
	db.DB.First(&paramedico, params["id"])

	if paramedico.ID == 0 {

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Paramedico no encontrado"))
		return
	}

	var updatedParamedico models.Paramedico
	json.NewDecoder(r.Body).Decode(&updatedParamedico)

	paramedico.DNI = updatedParamedico.DNI
	paramedico.NombreCompleto = updatedParamedico.NombreCompleto
	paramedico.Estado = updatedParamedico.Estado

	db.DB.Save(&paramedico)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paramedico)
}

func DeleteParamedicoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var paramedico models.Paramedico

	if err := db.DB.First(&paramedico, params["id"]).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("paramedico no encontrado"))
		return
	}

	// Cambiar el estado a "baja"
	paramedico.Estado = "baja" // O el valor correspondiente que uses en tu enumeración de TipoEstado

	// Guardar el cambio en la base de datos
	if err := db.DB.Save(&paramedico).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error al actualizar el paramedico"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("paramedico eliminado lógicamente"))
}
