package medicamentos

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]Medicamento, error)
	GetOne(ctx context.Context, id_med uint) (Medicamento, error)
	GetByNombre(ctx context.Context, nombre string) (Medicamento, error)
	Create(ctx context.Context, medicamento *Medicamento) error
	Update(ctx context.Context, id_med uint, medicamento Medicamento) error
	Delete(ctx context.Context, id_med uint) error
}
