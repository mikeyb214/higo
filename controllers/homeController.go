//Package controllers contains controllers for the web application.
package controllers

import (
	"html/template"
	"net/http"

	"github.com/mikeyb214/higo/models"
)

//HomeController contains action behaviors for the home page
type HomeController struct {
	Controller
}

//MakeHomeController initializes a MakeHomeController with a WikiRepository
func MakeHomeController(repo models.WikiRepository, tmpl *template.Template) *HomeController {
	c := new(HomeController)
	c.repo = repo
	c.template = tmpl

	return c
}

//Render handles the default action.
func (v *HomeController) Render(w http.ResponseWriter, r *http.Request) {
	//Check repo
	if v.repo == nil {
		http.Error(w, "Must initialize controller using MakeWikiController.", http.StatusInternalServerError)
		return
	}

	//Check path
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	wikiTitles, err := v.repo.ListWikiTitles()
	if v.repo != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	v.renderView(w, "index", &models.Page{PageTitle: "Home", WikiTitles: wikiTitles})
}
