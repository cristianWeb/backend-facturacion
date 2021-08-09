package medicamentos

//sistema para los medicamentos

type Medicamento struct {
	ID        uint   `json:"id_med,omitempty"`
	Nombre    string `json:"nombre,omitempty"`
	Precio    string `json:"precio,omitempty"`
	Ubicacion string `json:"ubicacion,omitempty"`
}

/*
Estructura basica para manejar la informacion de la base de datos
omitempty = omite el campo si esta vacio
*/
