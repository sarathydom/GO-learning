package render

import (
	"bytes"
	"exercise1/pkg/config"
	"exercise1/pkg/models"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var templateConfig *config.AppConfig
func NewTemplate(a *config.AppConfig) {
	templateConfig = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	// populate the template and cache them
	//templates,err:=CreateCacheTEmplate()
	var err error
	templates := templateConfig.TemplateCache


	// if err != nil {
	// 	log.Fatal(err)
	// }
	
	// get the cache template
	currentTemplate,ok := templates[tmpl]

	if !ok {
		log.Fatal(err)
	}
	// render the template

	buf:=new(bytes.Buffer)
	err=currentTemplate.Execute(buf,td)

	if err != nil { 
		log.Fatal(err)
	}

	_,err = buf.WriteTo(w)
	if err != nil { 
		log.Fatal(err)
	}
}

func CreateCacheTEmplate()(map[string]*template.Template,error) {
	var ts *template.Template
	var myTemplateCache = map[string]*template.Template{}

	filenames,err:=filepath.Glob("./templates/*.page.tmpl")

	if err != nil { 
		return myTemplateCache,err
	}

	for _,v:=range filenames {
		templateName := filepath.Base(v)
		ts,err = template.New(templateName).ParseFiles(v)
		if err != nil { 
			return myTemplateCache,err
		}
		layouts,err:=filepath.Glob("./templates/*.layout.tmpl")
		if err != nil { 
			return myTemplateCache,err
		}
		if len(layouts) > 0 {
			ts, err= ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil { 
				return myTemplateCache,err
			}
		}
		myTemplateCache[templateName] = ts
	}

	return myTemplateCache,nil
}


