package conexion_sql

import (
	//librerias PC
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	// Carpeta estructura
	"github.com/CarlosMarioM/Bia_energy/estructura"
)

// FUNCION PARA CONECTAR LA BD con GO
func ConectarBD() (*sql.DB, error) {

	//1. Conexion con parametros de ingreso a la BD
	const conexionStr = "user=postgres password=carlosmario12345 dbname=bia_energy sslmode=disable"

	//2. cConfirmacion de parametros y preparacion del gestor *sql (puntero)
	db, err := sql.Open("postgres", conexionStr)
	if err != nil {
		return nil, err
	}

	//3. Confirmacion de conexion fisica entre BD y GO
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	//Mensaje informativo
	fmt.Println("¡Conexión exitosa a la base de datos PostgreSQL!")
	return db, nil
}

// FUNCION PARA CONSULTAR REGISTROS E INSERTAR EN TABLA FINAL
func ObtenerTodosLosRegistros(db *sql.DB) ([]estructura.RegistroEnergia, error) {

	// 1. Se escreibe la consulta SELECT con los campos a llamar
	query := "SELECT id_serial, address, meter_id, active_energy, reading_timestamp FROM energy_readings"

	// 2. Se inicia la consulta
	filas, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	// 3. Al finalizar consulta se cierran los registros de la BD
	defer filas.Close()

	// 3. Se crea tabla interna del tipo  RefistroEnergia
	var tablaRegistros []estructura.RegistroEnergia

	/// 4. Se itera la estructura registro por registro
	//usando la variable registro
	for filas.Next() {
		var registro estructura.RegistroEnergia

		// 5. Con el puntero (&) se mapean los registros directamente desde la estructura
		//hasta la tabla interna que es la direccion final
		err := filas.Scan(&registro.IdSerial, &registro.Address, &registro.MeterId, &registro.ActiveEnergy, &registro.ReadingTimestamp)
		if err != nil {
			log.Printf("Error leyendo una fila: %v", err) //Mensaje de error por si algun fallo
			continue
		}

		// 6. Con APPEND se inserta el registro a la tabla
		tablaRegistros = append(tablaRegistros, registro)
	}

	// 7. Verificación final de seguridad, que confirma si
	// el bucle se realizo con exito de inicio a fin
	if errorFilas := filas.Err(); errorFilas != nil {
		return nil, errorFilas
	}

	return tablaRegistros, nil
}
