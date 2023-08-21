package handlers

import (
	"exercise1/pkg/render"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("into home page")
	render.RenderTemplate(w, "home.page.tmpl")
} 

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}