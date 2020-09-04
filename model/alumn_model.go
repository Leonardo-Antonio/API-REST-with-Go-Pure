package model

// Alumn is alumno of the table -> tb_alumnos
type Alumn struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Age      uint8  `json:"age"`
	Dni      string `json:"dni"`
}
