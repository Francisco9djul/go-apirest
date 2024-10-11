package routes

import (
	"encoding/json"

	"net/http"

	"github.com/Francisco9djul/go-apirest/db"
	"github.com/Francisco9djul/go-apirest/models"
	"github.com/gorilla/mux"
)

func GetReportesHandler(w http.ResponseWriter, r *http.Request) { //GET
	var Reportes []models.Reporte
	db.DB.Find(&Reportes)
	json.NewEncoder(w).Encode(&Reportes)
}

func GetReporteHandler(w http.ResponseWriter, r *http.Request) { //GET {id}
	var reporte models.Reporte
	params := mux.Vars(r)
	db.DB.First(&reporte, params["id"])

	if reporte.ID == 0 {

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Reporte no encontrado"))
		return

	}

	json.NewEncoder(w).Encode(&reporte)
}

func PostReporteHandler(w http.ResponseWriter, r *http.Request) { //POST
	var reporte models.Reporte

	json.NewDecoder(r.Body).Decode(&reporte)

	createdReporte := db.DB.Create(&reporte)
	err := createdReporte.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&reporte)
}

func PutReporteHandler(w http.ResponseWriter, r *http.Request) {
	var reporte models.Reporte
	params := mux.Vars(r)
	db.DB.First(&reporte, params["id"])

	if reporte.ID == 0 {

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Reporte no encontrado"))
		return
	}

	var updatedReporte models.Reporte
	json.NewDecoder(r.Body).Decode(&updatedReporte)

	reporte.Descripcion = updatedReporte.Descripcion
	reporte.FechayHora = updatedReporte.FechayHora

	db.DB.Save(&reporte)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reporte)
}

func DeleteReporteHandler(w http.ResponseWriter, r *http.Request) { //DELETE {id} (LOGICO)
	var reporte models.Reporte
	params := mux.Vars(r)
	db.DB.First(&reporte, params["id"])

	if reporte.ID == 0 {

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Reporte no encontrado"))
		return
	}

	db.DB.Delete(&reporte)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Reporte eliminado lógicamente"))
}
