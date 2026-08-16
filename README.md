# ⚡ BIA Energy Microservice - Backend API (Go)

Microservicio backend desarrollado en **Golang** para la gestión, procesamiento y consulta de series de tiempo de consumos energéticos de medidores. Este sistema está construido bajo los principios de **Clean Architecture**, asegurando escalabilidad, mantenibilidad y un rendimiento óptimo.

---

## 🚀 Características Principales

- **Conexión a Base de Datos Relacional:** Integración robusta con PostgreSQL para la lectura masiva de registros históricos de energía.
- **Motor de Series de Tiempo (Time-Series Engine):** Capacidad avanzada de filtrado y agrupación temporal orientada a cortes **Diarios (`daily`)**, **Semanales (`weekly`)** y **Mensuales (`monthly`)**.
- **Contratos JSON Exactos:** Generación automática de respuestas bajo el estándar de contratos de la API requerida, incluyendo relleno inteligente de ceros (`0`) para variables complementarias.
- **Pruebas Unitarias Integradas:** Módulo de pruebas (`_test.go`) para asegurar la integridad de las fórmulas matemáticas y lógicas.

---

## 🛠️ Tecnologías y Librerías Utilizadas

- **Lenguaje:** Go (Golang)
- **Base de Datos:** PostgreSQL
- **Driver SQL:** `github.com/lib/pq`
- **Manejo HTTP Nativo:** `net/http` (Alto rendimiento y bajo acoplamiento)

---

## 📂 Arquitectura del Proyecto

El proyecto sigue una estructura modular limpia:
```text
.
├── conexion_sql/       # Conexión y consultas masivas a PostgreSQL
├── consumos/           # Lógica de negocio, algoritmos de tiempo y reportes
├── estructura/         # Definición de estructuras (Structs) y modelos de datos
├── inicializar_programa/ # Archivo principal (main.go) y arranque del servidor web
└── tests/              # Pruebas automatizadas y unitarias