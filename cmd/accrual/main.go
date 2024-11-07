package main

import (
	"gophermart/internal/accrual/handler"
	"gophermart/internal/accrual/storage"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	s, err := setupStorage()
	if err != nil {
		panic(err)
	}

	r := setupRouter(s)

	err = http.ListenAndServe("localhost:8080", r)

	if err != nil {
		panic(err)
	}
}

func setupRouter(_ *storage.Storager) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Group(func(r chi.Router) {
		r.Use(middleware.Compress(9))
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/orders", func(r chi.Router) {
			r.Post("/", handler.OrderRegistrationHandler())
			r.Get("/{number}", handler.AccrualsCalculationHandler())
		})
		r.Route("/goods", func(r chi.Router) {
			r.Post("/", handler.RewardRegistrationHandler())
		})
	})

	return r
}

func setupStorage() (*storage.Storager, error) {
	repoConfig := &storage.Config{StorageType: storage.StorageTypeInmemory}

	storage, err := storage.NewStorage(repoConfig)
	if err != nil {
		return nil, err
	}

	return &storage, nil
}
