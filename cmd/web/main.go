package main

import (
	"exercise1/pkg/config"
	"exercise1/pkg/handlers"
	"exercise1/pkg/render"
	"log"
	"net/http"
)

func main() { 
	var app config.AppConfig
	template,err := render.CreateCacheTEmplate()
	if err != nil { 
		log.Fatal(err)
	}
	app.TemplateCache = template
	render.NewTemplate(&app)
	Repo := handlers.NewRepo(&app)
	handlers.NewHandlers(Repo)
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	// http.ListenAndServe(":2000", nil)

	srv := &http.Server{ 
		Addr: ":2000",
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}