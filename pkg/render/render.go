package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func RenderTemplate_old(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w,nil)
	if err != nil {
		fmt.Fprintf(w, "Error executing template")
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	log.Println("into render  func")
	var existingTemplate *template.Template
	var err error

	_, inMap := tc[tmpl]
	log.Println(inMap)
	if !inMap {
		log.Println("creating Cache template")
		err = createTemplateCache(tmpl)
	} else {
		log.Println("using Cache template")
	}

	existingTemplate = tc[tmpl]
	err = existingTemplate.Execute(w,nil)
	if err != nil { 
		log.Println(err)
	}

}

func createTemplateCache(tmpl string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s",tmpl),
		"./templates/base.layout.tmpl",
	}

	t,err := template.ParseFiles(templates...)
	if err != nil { 
		return err
	}
	tc[tmpl] = t
	return nil
}
