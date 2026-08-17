package consumos

import (
	"testing"

	"github.com/CarlosMarioM/Bia_energy/estructura"
)

// TestGenerarReporte valida que las operaciones y los filtros funcionen correctamente
func TestGenerarReporte(t *testing.T) {
	// 1. Datos simulados (Mock) para la prueba
	registrosMock := []estructura.RegistroEnergia{
		{
			MeterId:          "1",
			ActiveEnergy:     150.0,
			ReadingTimestamp: "2023-06-01T12:00:00Z",
		},
		{
			MeterId:          "1",
			ActiveEnergy:     50.0,
			ReadingTimestamp: "2023-06-02T12:00:00Z",
		},
	}

	// 2. Se ejecuta la función directamente
	resultado := GenerarReporte(registrosMock, "2023-06-01", "2023-06-02", "1", "daily")

	// 3. Se verifica la devolucion de resultados
	if len(resultado.Period) != 2 {
		t.Errorf("Se esperaban 2 periodos, se obtuvieron %d", len(resultado.Period))
	}

	// 4. Se verifica que la suma del primer día sea exacta (150.0)
	if resultado.DataGraph[0].Active[0] != 150.0 {
		t.Errorf("Error matemático: Se esperaba 150.0, se obtuvo %.2f", resultado.DataGraph[0].Active[0])
	}
}
