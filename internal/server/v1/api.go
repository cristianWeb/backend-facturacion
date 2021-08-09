package v1

import (
	"Sistema/internal/data"
	"net/http"

	"github.com/go-chi/chi"
)

func New() http.Handler {
	r := chi.NewRouter()

	ur := &MedicamentoRouter{
		Repository: &data.MedicamentoRepository{
			Data: data.New()},
	}

	r.Mount("/medicamentos", ur.Routes())

	return r
}
