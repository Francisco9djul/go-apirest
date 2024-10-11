package routes

import (
	"encoding/json"

	"net/http"

	"github.com/Francisco9djul/go-apirest/db"
	"github.com/Francisco9djul/go-apirest/models"
	"github.com/gorilla/mux"
)

func GetHospitalsHandler(w http.ResponseWriter, r *http.Request) { //GET
	var hospitals []models.Hospital

	// Cargar todos los hospitales junto con sus accidentes
	if err := db.DB.Preload("Accidentes").Find(&hospitals).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error al obtener los hospitals"))
		return
	}
	json.NewEncoder(w).Encode(&hospitals)
}

func GetHospitalHandler(w http.ResponseWriter, r *http.Request) { //GET {id}
	var hospital models.Hospital
	params := mux.Vars(r)

	if err := db.DB.Preload("Accidentes").First(&hospital, params["id"]).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("hospital no encontrado"))
		return
	}

	json.NewEncoder(w).Encode(&hospital)
}

func PostHospitalHandler(w http.ResponseWriter, r *http.Request) { //POST
	var hospital models.Hospital

	json.NewDecoder(r.Body).Decode(&hospital)

	createdHospital := db.DB.Create(&hospital)
	err := createdHospital.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&hospital)
}

func PutHospitalHandler(w http.ResponseWriter, r *http.Request) {
	var hospital models.Hospital
	params := mux.Vars(r)
	db.DB.First(&hospital, params["id"])

	if hospital.ID == 0 {

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Hospital no encontrado"))
		return
	}

	var updatedHospital models.Hospital
	json.NewDecoder(r.Body).Decode(&updatedHospital)

	hospital.Nombre = updatedHospital.Nombre
	hospital.Direccion = updatedHospital.Direccion

	db.DB.Save(&hospital)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hospital)
}

func DeleteHospitalHandler(w http.ResponseWriter, r *http.Request) { //DELETE {id} (LOGICO)
	var hospital models.Hospital
	params := mux.Vars(r)
	db.DB.First(&hospital, params["id"])

	if hospital.ID == 0 {

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Hospital no encontrado"))
		return
	}

	db.DB.Delete(&hospital)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hospital eliminado lógicamente"))
}
