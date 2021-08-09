package facturas

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]Factura, error)
	GetOne(ctx context.Context, id_factura uint) (Factura, error)
	GetByfecha_crear(ctx context.Context, fecha_crear string) (Factura, error)
	Create(ctx context.Context, factura *Factura) error
	Update(ctx context.Context, id_factura uint, factura Factura) error
	Delete(ctx context.Context, id_factura uint) error
}
