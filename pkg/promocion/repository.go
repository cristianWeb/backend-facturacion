package promocion

import "context"

//repositorio para el manejo del CRUD operacional de las promociones
type Repository interface {
	GetAll(ctx context.Context) ([]Promocion, error)
	GetOne(ctx context.Context, id_promocion uint) (Promocion, error)
	Create(ctx context.Context, promocion *Promocion) error
	Update(ctx context.Context, id_promocion uint, promocion Promocion) error
	Delete(ctx context.Context, id_promocion uint) error
}
