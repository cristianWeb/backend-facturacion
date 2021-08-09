package v1

import (
	"Sistema/pkg/medicamentos"
	"Sistema/pkg/response"
	"encoding/json"
	"net/http"
	"strconv"

	//"github.com/go-chi/chi"
	"fmt"

	"github.com/go-chi/chi"
)

type MedicamentoRouter struct {
	Repository medicamentos.Repository
}

func (ur *MedicamentoRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var u medicamentos.Medicamento
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = ur.Repository.Create(ctx, &u)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), u.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"medicamento": u})
}

//Funcion para traer todos  datos de los medicamentos

func (ur *MedicamentoRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	medicamentos, err := ur.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"medicamentos": medicamentos})
}

//Funcion para traer los datos de un medicamento

func (ur *MedicamentoRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id_med")

	id_med, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	u, err := ur.Repository.GetOne(ctx, uint(id_med))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"medicamento": u})
}

//Definiendo las rutas
func (ur *MedicamentoRouter) Routes() http.Handler {
	r := chi.NewRouter()
	// TODO: add routes.

	r.Get("/", ur.GetAllHandler)

	r.Post("/", ur.CreateHandler)

	r.Get("/{id_med}", ur.GetOneHandler)

	return r
}
