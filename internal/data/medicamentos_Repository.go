package data

import (
	"Sistema/pkg/medicamentos"
	"context"
)

type MedicamentoRepository struct {
	Data *Data
}

//Comencemos por el método GetAll, que nos permitirá obtener todos los registros en la tabla medicamento.
func (ur *MedicamentoRepository) GetAll(ctx context.Context) ([]medicamentos.Medicamento, error) {
	q := `
	SELECT id_med, nombre, precio, ubicacion
	FROM medicamento;
	`

	rows, err := ur.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var medicamentoss []medicamentos.Medicamento
	for rows.Next() {
		var u medicamentos.Medicamento
		rows.Scan(&u.ID, &u.Nombre, &u.Precio, &u.Ubicacion)
		medicamentoss = append(medicamentoss, u)
	}

	return medicamentoss, nil
}

func (ur *MedicamentoRepository) GetOne(ctx context.Context, id_med uint) (medicamentos.Medicamento, error) {
	q := `
    SELECT id_med, nombre, precio, ubicacion
        FROM medicamento WHERE id_med = $1;
    `

	row := ur.Data.DB.QueryRowContext(ctx, q, id_med)

	var u medicamentos.Medicamento
	err := row.Scan(&u.ID, &u.Nombre, &u.Precio, &u.Ubicacion)
	if err != nil {
		return medicamentos.Medicamento{}, err
	}

	return u, nil
}

func (ur *MedicamentoRepository) GetByNombre(ctx context.Context, Nombre string) (medicamentos.Medicamento, error) {
	q := `
	SELECT id_med, nombre, precio, ubicacion
	
        FROM medicamento WHERE Nombre = $1;
    `

	row := ur.Data.DB.QueryRowContext(ctx, q, Nombre)

	var u medicamentos.Medicamento
	err := row.Scan(&u.ID, &u.Nombre, &u.Precio, &u.Ubicacion)
	if err != nil {
		return medicamentos.Medicamento{}, err
	}

	return u, nil
}

func (ur *MedicamentoRepository) Create(ctx context.Context, u *medicamentos.Medicamento) error {
	q := `
    INSERT INTO medicamento (nombre, precio, ubicacion)
        VALUES ($_nombre, $_precio, $_ubicacion)
        RETURNING id_med;
    `

	row := ur.Data.DB.QueryRowContext(
		ctx, q, u.Nombre, u.Precio, u.Ubicacion)

	err := row.Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ur *MedicamentoRepository) Update(ctx context.Context, id_med uint, u medicamentos.Medicamento) error {
	q := `
    UPDATE medicamento set nombre=$_nombre, precio=$_precio, ubicacion=$_ubicacion
        WHERE id_med=$_id_med;
    `

	stmt, err := ur.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, u.Nombre, u.Precio, u.Ubicacion, id_med)
	if err != nil {
		return err
	}

	return nil
}
