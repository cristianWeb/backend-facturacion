package facturas

type Factura struct {
	ID         uint   `json:"id_factura,omitempty"`
	Fecha_cear string `json:"fecha_crea,omitempty"`
	Pago_total string `json:"pago_total,omitempty"`
}
