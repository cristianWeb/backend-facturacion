package promocion

//Post para crear un medicamento
type Promocion struct {
	ID           uint    `json:"id_promocion,omitempty"`
	Descripcion  string  `json:"descripcion,omitempty"`
	Porcentaje   float64 ` json:"porcentaje,omitempty"`
	Fecha_inicio string  `json:"fecha_inicio,omitempty"`
	Fecha_fin    string  `json:"fecha_fin,omitempty"`
}
