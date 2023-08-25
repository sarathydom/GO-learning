package handlers

import (
	"exercise1/pkg/config"
	"exercise1/pkg/models"
	"exercise1/pkg/render"
	"log"
	"net/http"
)




var Repo *Repository
type Repository struct {
	App *config.AppConfig
}



func NewRepo(app *config.AppConfig) *Repository {
	return &Repository{ 
		App : app,
	}
}

func NewHandlers(r *Repository) *Repository {
	repo := r
	return repo
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	log.Println("into home page")
	StringMap := make(map[string]string)
	StringMap["test"] = "HElloooo"
	
	render.RenderTemplate(w, "home.page.tmpl",&models.TemplateData{
		StringMap: StringMap,
	})
} 

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl",&models.TemplateData{})
}