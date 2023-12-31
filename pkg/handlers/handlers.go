package handlers

import (
	"github.com/theumee/golang-udemy/pkg/config"
	"github.com/theumee/golang-udemy/pkg/models"
	"github.com/theumee/golang-udemy/pkg/render"
	"net/http"
)

// Repo the repository used by our handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, Again!"

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{})

}
