package conexion_sql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	// Importamos la carpeta estructura para usar el molde
	"github.com/CarlosMarioM/Bia_energy/estructura"
)

func ConectarBD() (*sql.DB, error) {
	const conexionStr = "user=postgres password=carlosmario12345 dbname=bia_energy sslmode=disable"

	db, err := sql.Open("postgres", conexionStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("¡Conexión exitosa a la base de datos PostgreSQL!")
	return db, nil
}

// NUEVA FUNCIÓN
func ObtenerTodosLosRegistros(db *sql.DB) ([]estructura.RegistroEnergia, error) {
	// 1. Escribimos la consulta SQL
	query := "SELECT id_serial, address, meter_id, active_energy, reading_timestamp FROM energy_readings"

	// 2. Ejecutamos la consulta
	filas, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	// Cerramos las filas al terminar para no saturar la memoria
	defer filas.Close()

	// 3. Creamos nuestra "tabla interna" vacía
	var tablaRegistros []estructura.RegistroEnergia

	// 4. Recorremos fila por fila
	for filas.Next() {
		var registro estructura.RegistroEnergia

		// Empacamos los datos de PostgreSQL en nuestro molde de Go
		err := filas.Scan(&registro.IdSerial, &registro.Address, &registro.MeterId, &registro.ActiveEnergy, &registro.ReadingTimestamp)
		if err != nil {
			log.Printf("Error leyendo una fila: %v", err)
			continue
		}

		// Agregamos el registro lleno a la tabla interna
		tablaRegistros = append(tablaRegistros, registro)
	}

	// 5. Verificación final de seguridad (Solución exacta para el Linter)
	// Usamos una variable nueva "errorFilas" para que VS Code no se confunda
	if errorFilas := filas.Err(); errorFilas != nil {
		return nil, errorFilas
	}

	return tablaRegistros, nil
}
