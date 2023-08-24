package render

import (
	"bytes"
	"github.com/theumee/golang-udemy/pkg/config"
	"github.com/theumee/golang-udemy/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// NewTemplates sets templates using template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, html string, td *models.TemplateData) {

	// create a template cache
	tc := app.TemplateCache

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[html]
	if !ok {
		log.Fatal("could not get from template cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, td)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	//myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	//get all the files named *.page.html from ./templates

	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.html

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}

//var tc = make(map[string]*template.Template)
//
//func RenderTemplate(w http.ResponseWriter, t string) {
//	var tmpl *template.Template
//	var err error
//
//	//check to see if we already have the template in our cache
//	_, inMap := tc[t]
//	if !inMap {
//		//need to create the template
//		log.Println("creating template and adding to cache")
//		err = createTemplateCache(t)
//		if err != nil {
//			log.Println(err)
//		}
//	} else {
//		//we have the template in cache
//		log.Println("using cache template")
//	}
//
//	tmpl = tc[t]
//	err = tmpl.Execute(w, nil)
//}
//
//func createTemplateCache(t string) error {
//	templates := []string{
//		fmt.Sprintf("./templates/%s", t),
//		"./templates/base.layout.html",
//	}
//
//	tmpl, err := template.ParseFiles(templates...)
//	if err != nil {
//		return err
//	}
//
//	//add template to cache (map)
//
//	tc[t] = tmpl
//
//	return nil
//}
