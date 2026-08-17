# Documentación Técnica: Microservicio BIA Energy.

## 1.  Descripción General del Proyecto

Este proyecto es un microservicio backend desarrollado en **Golang** diseñado para el, procesamiento, agregación y consulta de series temporales de consumo energético. Cumple con los requerimientos técnicos para la prueba de **BIA Energy**.


## 2. Tecnologías y Stack

* **Lenguaje de Programación:** Go (Golang v1.18+)
* **Base de Datos Relacional:** PostgreSQL
* **Driver SQL:** `github.com/lib/pq`
* **Servidor HTTP:** Librería estándar `net/http` nativa de Go (optimizada para alta concurrencia).
* **Testing:** Librería estándar `testing` nativa de Go.

---

## 3. Arquitectura del Sistema (Clean Architecture)

El repositorio está estructurado en carpetas que son las siguientes:

```text
.
├── conexion_sql/       # [Capa de Infraestructura] Conexión a DB, consultas SQL y mapeo de datos.
├── consumos/           # [Capa de Casos de Uso] Reglas de negocio, agrupación temporal (daily, weekly, monthly) y test.
├── estructura/         # [Capa de Dominio] Modelos y contratos puros (Structs y JSON API).
└── main.go             # [Entrypoint] Inyección de dependencias, inicialización del servidor HTTP y validacion URL
```

---

## 4. Flujo

1.  **Ingesta de Petición HTTP:** El cliente envía una consulta GET al endpoint `/consumption` con los filtros de tiempo y medidores requeridos.
2.  **Capa de Validación:** El manejador en `main.go` verifica la presencia de los parámetros obligatorios implementando *Early Returns* para responder un `400 Bad Request` en caso de datos faltantes.
3.  **Procesamiento de Negocio:** La capa `consumos` recibe el histórico en memoria y ejecuta los algoritmos de segmentación temporal (búsqueda de inicio de semana, agrupación por meses y filtro por medidor).
4.  **Respuesta JSON:** Se codifica (serialize) el arreglo final mapeando la salida a la estructura de la API especificada en los requerimientos.

---

## 5. Guía para la instalación y ejecución local

### 5.1. Requisitos Previos
* Go instalado en su máquina local.
* Servidor PostgreSQL ejecutándose.
* Cargar el archivo `dataset.csv` proporcionado hacia una tabla llamada `energy_readings` en la base de datos PostgreSQL.

### 5.2. Configuración de Credenciales
Edite el archivo `conexion_sql/conexion.go` y modifique la constante `conexionStr` con sus credenciales locales:
```go
const conexionStr = "user=TU_USUARIO password=TU_PASSWORD dbname=bia_energy sslmode=disable"
```

### 5.3. Levantar el Servidor
```bash
# Descargar dependencias del driver de Postgres
go mod tidy

# Compilar y arrancar el servidor
go run main.go
```
*Salida esperada en consola:*
> `¡Conexión exitosa a la base de datos PostgreSQL!`
> `Datos cargados en memoria exitosamente.`
> `Servidor ejecutado en => http://localhost:8080`

---

## 6. Referencia de la API REST

### Endpoint Principal
`GET /consumption`

Este endpoint permite consultar la gráfica de consumos de energía filtrando por rangos temporales específicos y frecuencias.

### Parámetros de Consulta Obligatorios (Query Parameters)
| Parámetro | Tipo | Descripción | Ejemplo |
| :--- | :--- | :--- | :--- |
| `meters_ids` | `string` | IDs de los medidores a consultar, separados por comas. | `1,2,3` |
| `start_date` | `string` | Límite inferior del rango de fechas (Formato `YYYY-MM-DD`). | `2022-08-01` |
| `end_date` | `string` | Límite superior del rango de fechas (Formato `YYYY-MM-DD`). | `2022-08-31` |
| `kind_period` | `string` | Frecuencia de la agregación de los datos. Valores permitidos: `daily`, `weekly`, `monthly`. | `monthly` |

---

## 7. Ejemplos de Peticiones y Respuestas (cURL)

### 7.1. Frecuencia Mensual (Monthly)
Agrupa las lecturas de cada medidor por mes calendario.

**Petición:**
```bash
curl -X GET "http://localhost:8080/consumption?meters_ids=1,2&start_date=2022-08-01&end_date=2022-10-31&kind_period=monthly"
```

**Respuesta Exitosa (200 OK):**
```json
{
  "period": [
    "2022-08-01",
    "2022-09-01",
    "2022-10-01"
  ],
  "data_graph": [
    {
      "meter_id": "1",
      "address": "Direccion Mock",
      "data_data": [ 1500.5, 1420.0, 1600.2 ]
    }
  ]
}
```

### 7.2. Frecuencia Semanal (Weekly)
Agrupa las lecturas calculando el inicio correcto de cada semana (Lunes).

**Petición:**
```bash
curl -X GET "http://localhost:8080/consumption?meters_ids=1,2&start_date=2022-08-01&end_date=2022-08-15&kind_period=weekly"
```

### 7.3. Frecuencia Diaria (Daily)
Retorna la lectura diaria.

**Petición:**
```bash
curl -X GET "http://localhost:8080/consumption?meters_ids=1,2,3&start_date=2022-08-01&end_date=2022-08-03&kind_period=daily"
```