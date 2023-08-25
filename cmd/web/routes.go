package main

import (
	"exercise1/pkg/config"
	"exercise1/pkg/handlers"
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/go-chi/chi/v5"
)

func routes_pat(app *config.AppConfig) http.Handler{
	mux:=pat.New()
	mux.Get("/",http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about",http.HandlerFunc(handlers.Repo.About))
	return mux
}

func routes(any *config.AppConfig) http.Handler {
	mux:=chi.NewRouter()
	mux.Use(console)
	mux.Use(nosur)
	mux.Get("/",handlers.Repo.Home)
	mux.Get("/about",handlers.Repo.About)
	return mux
}