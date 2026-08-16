-- 1. Creacion de tabla en pg
CREATE TABLE IF NOT EXISTS energy_readings (

-- CAMPOS:

--2. Identificador de pg con id de autoincremento por cada registro
id_serial SERIAL PRIMARY KEY,

--3. Creacion de campo meter_id que va a contener las direccioness de cada medidor fisico
--(NOT NULL) para volverlo obligatorio, de no registrarlo sera rechazado el registro
address VARCHAR(100) NOT NULL,


--4. Creacion de id de medidor (va del 1 al 3)
meter_id INT NOT NULL,


--5. Creacion de campo active_energy de tipo decimal para registrar las lecturas
-- 12 digitos enteros x 4 decimales, tambien obligatorio
    active_energy NUMERIC(12, 4) NOT NULL,

--6. Creacion de campo reading_timestamp para registrar fecha y hora de la lectura del medidor
--guardara fecha y hora exacta (AAAA/MM/DD) y (HH:MM:SS)
reading_timestamp TIMESTAMP NOT NULL
);