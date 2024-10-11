package routes

import (
	"encoding/json"

	"net/http"

	"github.com/Francisco9djul/go-apirest/db"
	"github.com/Francisco9djul/go-apirest/models"
	"github.com/gorilla/mux"
)

func GetPacientesHandler(w http.ResponseWriter, r *http.Request) { //GET
	var pacientes []models.Paciente

	// Cargar todos los pacientes junto con sus accidentes
	if err := db.DB.Preload("Accidentes").Find(&pacientes).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error al obtener los pacientes"))
		return
	}
	json.NewEncoder(w).Encode(&pacientes)
}

func GetPacienteHandler(w http.ResponseWriter, r *http.Request) { //GET {id}
	var paciente models.Paciente
	params := mux.Vars(r)

	if err := db.DB.Preload("Accidentes").First(&paciente, params["id"]).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("paciente no encontrado"))
		return
	}

	json.NewEncoder(w).Encode(&paciente)
}

func PostPacienteHandler(w http.ResponseWriter, r *http.Request) { //POST
	var paciente models.Paciente

	json.NewDecoder(r.Body).Decode(&paciente)

	createdPaciente := db.DB.Create(&paciente)
	err := createdPaciente.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&paciente)
}

func PutPacienteHandler(w http.ResponseWriter, r *http.Request) {
	var paciente models.Paciente
	params := mux.Vars(r)
	db.DB.First(&paciente, params["id"])

	if paciente.ID == 0 {

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("paciente no encontrado"))
		return
	}

	var updatedPaciente models.Paciente
	json.NewDecoder(r.Body).Decode(&updatedPaciente)

	paciente.DNI = updatedPaciente.DNI
	paciente.NombreCompleto = updatedPaciente.NombreCompleto
	paciente.Telefono = updatedPaciente.Telefono

	db.DB.Save(&paciente)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paciente)
}

func DeletePacienteHandler(w http.ResponseWriter, r *http.Request) { //DELETE {id} (LOGICO)
	var paciente models.Paciente
	params := mux.Vars(r)
	db.DB.First(&paciente, params["id"])

	if paciente.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Paciente no encontrado"))
		return
	}

	db.DB.Delete(&paciente)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Paciente eliminado lógicamente"))
}
