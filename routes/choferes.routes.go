package routes

import (
	"encoding/json"

	"net/http"

	"github.com/Francisco9djul/go-apirest/db"
	"github.com/Francisco9djul/go-apirest/models"
	"github.com/gorilla/mux"
)

func GetChoferesHandler(w http.ResponseWriter, r *http.Request) { //GET
	var choferes []models.Chofer

	// Cargar todos los choferes junto con sus ambulancias
	if err := db.DB.Preload("Ambulancias").Find(&choferes).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error al obtener los choferes"))
		return
	}

	json.NewEncoder(w).Encode(&choferes)
}

func GetChoferHandler(w http.ResponseWriter, r *http.Request) { //GET {id}
	var chofer models.Chofer
	params := mux.Vars(r)

	if err := db.DB.Preload("Ambulancias").First(&chofer, params["id"]).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Chofer no encontrado"))
		return
	}

	json.NewEncoder(w).Encode(&chofer)
}

func PostChoferHandler(w http.ResponseWriter, r *http.Request) { //POST
	var chofer models.Chofer

	json.NewDecoder(r.Body).Decode(&chofer)

	createdChofer := db.DB.Create(&chofer)
	err := createdChofer.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&chofer)
}

func PutChoferHandler(w http.ResponseWriter, r *http.Request) {
	var chofer models.Chofer
	params := mux.Vars(r)
	db.DB.First(&chofer, params["id"])

	if chofer.ID == 0 {

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Chofer no encontrado"))
		return
	}

	var updatedChofer models.Chofer
	json.NewDecoder(r.Body).Decode(&updatedChofer)

	chofer.DNI = updatedChofer.DNI
	chofer.NombreCompleto = updatedChofer.NombreCompleto
	chofer.Estado = updatedChofer.Estado

	db.DB.Save(&chofer)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chofer)
}

func DeleteChoferHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var chofer models.Chofer
	// Buscar el chofer en la base de datos
	if err := db.DB.First(&chofer, params["id"]).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Chofer no encontrado"))
		return
	}

	// Cambiar el estado a "baja"
	chofer.Estado = "baja" // O el valor correspondiente que uses en tu enumeración de TipoEstado

	// Guardar el cambio en la base de datos
	if err := db.DB.Save(&chofer).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error al actualizar el chofer"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Chofer eliminado lógicamente"))
}
