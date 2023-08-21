package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	// populate the template and cache them
	templates,err:=createCacheTEmplate()
	if err != nil {
		log.Fatal(err)
	}
	
	// get the cache template
	currentTemplate,ok := templates[tmpl]

	if !ok {
		log.Fatal(err)
	}
	// render the template

	buf:=new(bytes.Buffer)
	err=currentTemplate.Execute(buf,nil)

	if err != nil { 
		log.Fatal(err)
	}

	_,err = buf.WriteTo(w)
	if err != nil { 
		log.Fatal(err)
	}
}

func createCacheTEmplate()(map[string]*template.Template,error) {
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


