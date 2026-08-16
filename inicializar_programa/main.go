package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/CarlosMarioM/Bia_energy/conexion_sql"
	"github.com/CarlosMarioM/Bia_energy/consumos"
)

func main() {
	fmt.Println("Iniciando el servidor API de BIA Energy...")

	db, err := conexion_sql.ConectarBD()
	if err != nil {
		log.Fatalf("Error crítico de BD: %v", err)
	}

	tablaRegistros, err := conexion_sql.ObtenerTodosLosRegistros(db)
	if err != nil {
		log.Fatalf("Error trayendo registros: %v", err)
	}
	fmt.Println("Datos cargados en memoria exitosamente.")

	http.HandleFunc("/consumption", func(w http.ResponseWriter, r *http.Request) {
		medidores := r.URL.Query().Get("meters_ids")
		fechaInicio := r.URL.Query().Get("start_date")
		fechaFin := r.URL.Query().Get("end_date")
		kindPeriod := r.URL.Query().Get("kind_period")

		if medidores == "" || fechaInicio == "" || fechaFin == "" || kindPeriod == "" {
			http.Error(w, "Faltan parámetros en la URL", http.StatusBadRequest)
			return
		}

		cajaFinal := consumos.GenerarReporte(tablaRegistros, fechaInicio, fechaFin, medidores, kindPeriod)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cajaFinal)
	})

	fmt.Println("✅ Servidor corriendo y escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
