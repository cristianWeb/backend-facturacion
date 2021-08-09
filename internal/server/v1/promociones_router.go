package v1

import (
	"Sistema/pkg/promocion"
	"Sistema/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type PromocionRouter struct {
	Repository promocion.Repository
}

// crear promocion

func (pr *PromocionRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var p promocion.Promocion
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = pr.Repository.Create(ctx, &p)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), p.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"promocion": p})
}

//obtener promociones

func (pr *PromocionRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	promociones, err := pr.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"promocion": promociones})
}

//obtener una promocion

func (pr *PromocionRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id_promo, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	p, err := pr.Repository.GetOne(ctx, uint(id_promo))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"promocion": p})
}

//rutas de promociones

func (pr *PromocionRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", pr.GetAllHandler)

	r.Post("/", pr.CreateHandler)

	r.Get("/{id_promo}", pr.GetOneHandler)

	return r
}
