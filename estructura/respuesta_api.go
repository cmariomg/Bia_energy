package estructura

// ConsumoMedidor representa la estructura de energías de un solo medidor
type ConsumoMedidor struct {
	MeterID            int       `json:"meter_id"`
	Address            string    `json:"address"`
	Active             []float64 `json:"active"`
	ReactiveInductive  []float64 `json:"reactive_inductive"`
	ReactiveCapacitive []float64 `json:"reactive_capacitive"`
	Exported           []float64 `json:"exported"`
}

// RespuestaConsumo es la "Caja final" que enviaremos a internet
type RespuestaConsumo struct {
	Period    []string         `json:"period"`
	DataGraph []ConsumoMedidor `json:"data_graph"`
}
