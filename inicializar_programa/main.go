package main

import (
	//Librerias del sistema
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//Modulos propios, llaman las funciones
	"github.com/CarlosMarioM/Bia_energy/conexion_sql"
	"github.com/CarlosMarioM/Bia_energy/consumos"
)

func main() {

	//1. valida la conexion exitosa a la BD, y de no ser asi
	//   corta el programa con fatalf y manda el mensaje
	fmt.Println("Iniciando el servidor Bia Energy.")
	db, err := conexion_sql.ConectarBD()
	if err != nil {
		log.Fatalf("Error crítico de BD: %v", err)
	}

	//2. Se obtienen los registros
	tablaRegistros, err := conexion_sql.ObtenerTodosLosRegistros(db)
	if err != nil {
		log.Fatalf("Error trayendo registros: %v", err)
	}
	fmt.Println("Datos cargados en memoria exitosamentee.")

	//3. Se consultan los parametros de la URL (GET)
	//  y se asignan a las variables nuevas
	http.HandleFunc("/consumption", func(w http.ResponseWriter, r *http.Request) {
		medidores := r.URL.Query().Get("meters_ids")
		fechaInicio := r.URL.Query().Get("start_date")
		fechaFin := r.URL.Query().Get("end_date")
		kindPeriod := r.URL.Query().Get("kind_period")

		//4. Se comprueban que todas las variables esten llenas para seguir el programa
		if medidores == "" || fechaInicio == "" || fechaFin == "" || kindPeriod == "" {
			http.Error(w, "Faltan parámetros en la URL", http.StatusBadRequest)
			return
		}

		//5. CaajaFinal presenta el valor de las variables en estructura para que
		//	la funcion de consumos procese la info y devuelva los registros dado el requerimiento
		cajaFinal := consumos.GenerarReporte(tablaRegistros, fechaInicio, fechaFin, medidores, kindPeriod)

		//6.  La informacion es devuelta en formato JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cajaFinal)
	})

	//7. mensaje de ejecucion exitosa, y de no serla, FATAL
	// lo detiene
	fmt.Println(" Servidor ejecutado en => http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
