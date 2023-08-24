package main

import (
	"fmt"
	"github.com/theumee/golang-udemy/pkg/config"
	"github.com/theumee/golang-udemy/pkg/handlers"
	"github.com/theumee/golang-udemy/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting application on port: ", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
