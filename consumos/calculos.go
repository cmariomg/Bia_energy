package consumos

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/CarlosMarioM/Bia_energy/estructura"
)

const formatoBD = "2006-01-02T15:04:05Z"
const formatoAPI = "2006-01-02"

func obtenerMesCorto(mes time.Month) string {
	meses := map[time.Month]string{
		time.January: "JAN", time.February: "FEB", time.March: "MAR",
		time.April: "APR", time.May: "MAY", time.June: "JUN",
		time.July: "JUL", time.August: "AUG", time.September: "SEP",
		time.October: "OCT", time.November: "NOV", time.December: "DEC",
	}
	return meses[mes]
}

type RangoTiempo struct {
	Inicio time.Time
	Fin    time.Time
}

func GenerarReporte(registros []estructura.RegistroEnergia, startDate string, endDate string, metersIDsParam string, kindPeriod string) estructura.RespuestaConsumo {
	inicio, _ := time.Parse(formatoAPI, startDate)
	finGlobal, _ := time.Parse(formatoAPI, endDate)
	finGlobal = finGlobal.Add((24 * time.Hour) - time.Second)

	var periodos []string
	var rangos []RangoTiempo
	fechaActual := inicio

	// 1. CREAR EL EJE TEMPORAL (PERIOD)
	for fechaActual.Before(finGlobal) || fechaActual.Equal(finGlobal) {
		switch kindPeriod {
		case "monthly":
			inicioMes := time.Date(fechaActual.Year(), fechaActual.Month(), 1, 0, 0, 0, 0, time.UTC)
			finMes := inicioMes.AddDate(0, 1, -1).Add((24 * time.Hour) - time.Second)
			etiqueta := fmt.Sprintf("%s %d", obtenerMesCorto(inicioMes.Month()), inicioMes.Year())
			periodos = append(periodos, etiqueta)
			rangos = append(rangos, RangoTiempo{Inicio: inicioMes, Fin: finMes})
			fechaActual = inicioMes.AddDate(0, 1, 0)

		case "weekly":
			finSemana := fechaActual.AddDate(0, 0, 6)
			finRango := finSemana.Add((24 * time.Hour) - time.Second)
			etiqueta := fmt.Sprintf("%s %d - %s %d", obtenerMesCorto(fechaActual.Month()), fechaActual.Day(), obtenerMesCorto(finSemana.Month()), finSemana.Day())
			periodos = append(periodos, etiqueta)
			rangos = append(rangos, RangoTiempo{Inicio: fechaActual, Fin: finRango})
			fechaActual = fechaActual.AddDate(0, 0, 7)

		default: // "daily"
			finDia := fechaActual.Add((24 * time.Hour) - time.Second)
			etiqueta := fmt.Sprintf("%s %d", obtenerMesCorto(fechaActual.Month()), fechaActual.Day())
			periodos = append(periodos, etiqueta)
			rangos = append(rangos, RangoTiempo{Inicio: fechaActual, Fin: finDia})
			fechaActual = fechaActual.AddDate(0, 0, 1)
		}
	}

	cantidadCajones := len(periodos)

	// 2. SEPARAR LOS IDs DE LOS MEDIDORES POR COMAS (Ej: "1,2" -> ["1", "2"])
	listaIDs := strings.Split(metersIDsParam, ",")
	var listaDatosMedidores []estructura.ConsumoMedidor

	// 3. PROCESAR CADA MEDIDOR INDEPENDIENTEMENTE
	for _, mIDStr := range listaIDs {
		mIDTrim := strings.TrimSpace(mIDStr)
		meterIdInt, _ := strconv.Atoi(mIDTrim)

		datosMedidor := estructura.ConsumoMedidor{
			MeterID:            meterIdInt,
			Address:            fmt.Sprintf("Dirección del medidor %s", mIDTrim),
			Active:             make([]float64, cantidadCajones),
			ReactiveInductive:  make([]float64, cantidadCajones),
			ReactiveCapacitive: make([]float64, cantidadCajones),
			Exported:           make([]float64, cantidadCajones),
		}

		// Rellenar las energías de este medidor en su respectivo cajón temporal
		for _, r := range registros {
			if r.MeterId != mIDTrim {
				continue
			}

			fechaRegistro, err := time.Parse(formatoBD, r.ReadingTimestamp)
			if err != nil {
				continue
			}

			if (fechaRegistro.After(inicio) || fechaRegistro.Equal(inicio)) &&
				(fechaRegistro.Before(finGlobal) || fechaRegistro.Equal(finGlobal)) {

				for i, rango := range rangos {
					if (fechaRegistro.After(rango.Inicio) || fechaRegistro.Equal(rango.Inicio)) &&
						(fechaRegistro.Before(rango.Fin) || fechaRegistro.Equal(rango.Fin)) {
						datosMedidor.Active[i] += r.ActiveEnergy
						break
					}
				}
			}
		}

		listaDatosMedidores = append(listaDatosMedidores, datosMedidor)
	}

	// 4. EMPACAR TODO EN LA CAJA FINAL
	respuestaFinal := estructura.RespuestaConsumo{
		Period:    periodos,
		DataGraph: listaDatosMedidores,
	}

	return respuestaFinal
}
