package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/levshindenis/Auth_JWT-GO/internal/app/handlers"
)

func Router(mh handlers.MyHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Get("/access/{uuid}", mh.Access)
		r.Get("/refresh", mh.Refresh)
	})
	return r
}
