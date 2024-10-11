package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Francisco9djul/go-apirest/db"
	"github.com/Francisco9djul/go-apirest/models"
	"github.com/Francisco9djul/go-apirest/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection() // me conecto a la BBDD

	//Migraciones
	db.DB.AutoMigrate(models.Chofer{})
	db.DB.AutoMigrate(models.Paramedico{})
	db.DB.AutoMigrate(models.Ambulancia{})
	db.DB.AutoMigrate(models.Hospital{})
	db.DB.AutoMigrate(models.Reporte{})
	db.DB.AutoMigrate(models.Paciente{})
	db.DB.AutoMigrate(models.Accidente{})

	r := mux.NewRouter() // Creo el enrutador

	r.HandleFunc("/", routes.HomeHandler) // Defino la ruta principal

	//Llamada a METODOS de la clase HOSPITAL
	r.HandleFunc("/hospitals", routes.GetHospitalsHandler).Methods("GET")
	r.HandleFunc("/hospitals/{id}", routes.GetHospitalHandler).Methods("GET")
	r.HandleFunc("/hospitals", routes.PostHospitalHandler).Methods("POST")
	r.HandleFunc("/hospitals/{id}", routes.PutHospitalHandler).Methods("PUT")
	r.HandleFunc("/hospitals/{id}", routes.DeleteHospitalHandler).Methods("DELETE") //eliminado LOGICO (no fisico)

	//Llamada a METODOS de la clase REPORTE
	r.HandleFunc("/reportes", routes.GetReportesHandler).Methods("GET")
	r.HandleFunc("/reportes/{id}", routes.GetReporteHandler).Methods("GET")
	r.HandleFunc("/reportes", routes.PostReporteHandler).Methods("POST")
	r.HandleFunc("/reportes/{id}", routes.PutReporteHandler).Methods("PUT")
	r.HandleFunc("/reportes/{id}", routes.DeleteReporteHandler).Methods("DELETE") //eliminado LOGICO (no fisico)

	//Llamada a METODOS de la clase CHOFER
	r.HandleFunc("/choferes", routes.GetChoferesHandler).Methods("GET")
	r.HandleFunc("/choferes/{id}", routes.GetChoferHandler).Methods("GET")
	r.HandleFunc("/choferes", routes.PostChoferHandler).Methods("POST")
	r.HandleFunc("/choferes/{id}", routes.PutChoferHandler).Methods("PUT")
	r.HandleFunc("/choferes/{id}", routes.DeleteChoferHandler).Methods("DELETE") //eliminado LOGICO (no fisico)(solo cambio el estado)

	//Llamada a METODOS de la clase PARAMEDICO
	r.HandleFunc("/paramedicos", routes.GetParamedicosHandler).Methods("GET")
	r.HandleFunc("/paramedicos/{id}", routes.GetParamedicoHandler).Methods("GET")
	r.HandleFunc("/paramedicos", routes.PostParamedicoHandler).Methods("POST")
	r.HandleFunc("/paramedicos/{id}", routes.PutParamedicoHandler).Methods("PUT")
	r.HandleFunc("/paramedicos/{id}", routes.DeleteParamedicoHandler).Methods("DELETE") //eliminado LOGICO (no fisico)(solo cambio el estado)

	//Llamada a METODOS de la clase PACIENTE
	r.HandleFunc("/pacientes", routes.GetPacientesHandler).Methods("GET")
	r.HandleFunc("/pacientes/{id}", routes.GetPacienteHandler).Methods("GET")
	r.HandleFunc("/pacientes", routes.PostPacienteHandler).Methods("POST")
	r.HandleFunc("/pacientes/{id}", routes.PutPacienteHandler).Methods("PUT")
	r.HandleFunc("/pacientes/{id}", routes.DeletePacienteHandler).Methods("DELETE") //eliminado LOGICO (no fisico)

	//Llamada a METODOS de la clase AMBULANCIA
	r.HandleFunc("/ambulancias", routes.GetAmbulanciasHandler).Methods("GET")
	r.HandleFunc("/ambulancias/{id}", routes.GetAmbulanciaHandler).Methods("GET")
	r.HandleFunc("/ambulancias", routes.PostAmbulanciaHandler).Methods("POST")
	r.HandleFunc("/ambulancias/{id}", routes.PutAmbulanciaHandler).Methods("PUT")
	r.HandleFunc("/ambulancias/{id}", routes.DeleteAmbulanciaHandler).Methods("DELETE") //eliminado LOGICO (no fisico)

	//Llamada a METODOS de la clase AMBULANCIA
	r.HandleFunc("/accidentes", routes.GetAccidentesHandler).Methods("GET")
	r.HandleFunc("/accidentes/{id}", routes.GetAccidenteHandler).Methods("GET")
	r.HandleFunc("/accidentes", routes.PostAccidenteHandler).Methods("POST")
	r.HandleFunc("/accidentes/{id}", routes.PutAccidenteHandler).Methods("PUT")
	r.HandleFunc("/accidentes/{id}", routes.DeleteAccidenteHandler).Methods("DELETE") //eliminado LOGICO (no fisico)

	fmt.Println("Servidor iniciado en http://localhost:3000")

	log.Fatal(http.ListenAndServe(":3000", r)) // Inicializo el servidor HTTP con el enrutador
}
