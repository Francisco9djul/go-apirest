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

func GetAccidentesHandler(w http.ResponseWriter, r *http.Request) { //GET
	var accidentes []models.Accidente
	db.DB.Find(&accidentes)
	json.NewEncoder(w).Encode(&accidentes)
}

func GetAccidenteHandler(w http.ResponseWriter, r *http.Request) { //GET {id}
	var accidente models.Accidente
	params := mux.Vars(r)
	db.DB.First(&accidente, params["id"])

	if accidente.ID == 0 {

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("accidente no encontrado"))
		return

	}

	json.NewEncoder(w).Encode(&accidente)
}

func PostAccidenteHandler(w http.ResponseWriter, r *http.Request) {
	var accidente models.Accidente

	// Decodificar el cuerpo del request
	if err := json.NewDecoder(r.Body).Decode(&accidente); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error al decodificar el JSON"))
		return
	}

	// Verificar que las FK existan (pero permitir que se reutilicen)
	if !validarFk(w, accidente) {
		return
	}

	// Crear el nuevo accidente
	if err := db.DB.Create(&accidente).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error al crear el accidente: " + err.Error()))
		return
	}

	// Devolver el accidente creado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&accidente)
}

func validarFk(w http.ResponseWriter, accidente models.Accidente) bool {
	// Verificar si el Paciente relacionado existe
	var pacienteExistente models.Paciente
	if err := db.DB.First(&pacienteExistente, accidente.PacienteID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Paciente no encontrado"))
			return false
		}
	}

	// Verificar si la ambulancia relacionada existe
	var ambulanciaExistente models.Ambulancia
	if err := db.DB.First(&ambulanciaExistente, accidente.AmbulanciaID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Ambulancia no encontrada"))
			return false
		}
	}

	// Verificar si el hospital relacionado existe
	var hospitalExistente models.Hospital
	if err := db.DB.First(&hospitalExistente, accidente.HospitalID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Hospital no encontrado"))
			return false
		}
	}

	return true
}

//func PostAccidenteHandler(w http.ResponseWriter, r *http.Request) { //POST
//	var accidente models.Accidente
//
//	json.NewDecoder(r.Body).Decode(&accidente)
//
//	// Verificar si el Paciente relacionado existe
//	var pacienteExistente models.Paciente
//	if err := db.DB.First(&pacienteExistente, accidente.PacienteID).Error; err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			w.WriteHeader(http.StatusBadRequest) // 400
//			w.Write([]byte("Paciente no encontrado"))
//			return
//		}
//	}
//
//	// Verificar si el ambulancia relacionado existe
//	var ambulanciaExistente models.Ambulancia
//	if err := db.DB.First(&ambulanciaExistente, accidente.AmbulanciaID).Error; err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			w.WriteHeader(http.StatusBadRequest) // 400
//			w.Write([]byte("Ambulancia no encontrado"))
//			return
//		}
//	}
//
//	createdAccidente := db.DB.Create(&accidente)
//	err := createdAccidente.Error
//
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest) // 400
//		w.Write([]byte(err.Error()))
//	}
//
//	json.NewEncoder(w).Encode(&accidente)
//}

func PutAccidenteHandler(w http.ResponseWriter, r *http.Request) {
	var accidente models.Accidente
	params := mux.Vars(r)
	db.DB.First(&accidente, params["id"])

	if accidente.ID == 0 {

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("accidente no encontrado"))
		return
	}

	var updatedAccidente models.Accidente
	json.NewDecoder(r.Body).Decode(&updatedAccidente)

	// Verificar si el Paciente relacionado existe
	var pacienteExistente models.Paciente
	if err := db.DB.First(&pacienteExistente, accidente.PacienteID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusBadRequest) // 400
			w.Write([]byte("Paciente no encontrado"))
			return
		}
	}

	// Verificar si el ambulancia relacionado existe
	var ambulanciaExistente models.Ambulancia
	if err := db.DB.First(&ambulanciaExistente, accidente.AmbulanciaID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusBadRequest) // 400
			w.Write([]byte("Ambulancia no encontrado"))
			return
		}
	}

	accidente.Direccion = updatedAccidente.Direccion
	accidente.Descripcion = updatedAccidente.Descripcion
	accidente.RequiereTranslado = updatedAccidente.RequiereTranslado
	accidente.Fecha = updatedAccidente.Fecha
	accidente.Hora = updatedAccidente.Hora
	accidente.PacienteID = updatedAccidente.PacienteID
	accidente.AmbulanciaID = updatedAccidente.AmbulanciaID
	accidente.ReporteID = updatedAccidente.ReporteID
	accidente.HospitalID = updatedAccidente.HospitalID

	db.DB.Save(&accidente)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accidente)
}

func DeleteAccidenteHandler(w http.ResponseWriter, r *http.Request) { //DELETE {id} (LOGICO)
	var accidente models.Accidente
	params := mux.Vars(r)
	db.DB.First(&accidente, params["id"])

	if accidente.ID == 0 {

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Accidente no encontrado"))
		return
	}

	db.DB.Delete(&accidente)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Accidente eliminado lógicamente"))
}
