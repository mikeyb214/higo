//Package controllers contains controllers for the web application.
package controllers

import (
	"html/template"
	"net/http"
	"regexp"

	"github.com/mikeyb214/higo/models"
)

//WikiController contains action behaviors for Wikis
type WikiController struct {
	Controller
	validPath *regexp.Regexp
}

//MakeWikiController initializes a WikiController with a WikiRepository
func MakeWikiController(repo models.WikiRepository, tmpl *template.Template) *WikiController {
	c := new(WikiController)
	c.repo = repo
	c.template = tmpl
	c.validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

	return c
}

//Render handles the View action.
func (v *WikiController) Render(w http.ResponseWriter, r *http.Request) {
	//Check repo
	if v.repo == nil {
		http.Error(w, "Must initialize controller using MakeWikiController.", http.StatusInternalServerError)
		return
	}

	//Check path
	match := v.validPath.FindStringSubmatch(r.URL.Path)
	if match == nil {
		http.NotFound(w, r)
		return
	}
	title := match[2]

	//Do View
	wiki, err := v.repo.LoadWiki(title)
	if err != nil {
		//Did not find page, so switch to edit mode for new page
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	wikiTitles, err := v.repo.ListWikiTitles()
	if v.repo != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	v.renderView(w, "view", &models.Page{PageTitle: title, WikiTitles: wikiTitles, Wiki: wiki})
}

//Save handles the Save action
func (v *WikiController) Save(w http.ResponseWriter, r *http.Request) {
	//Check repo
	if v.repo == nil {
		http.Error(w, "Must initialize controller using MakeWikiController.", http.StatusInternalServerError)
		return
	}

	//Check path
	match := v.validPath.FindStringSubmatch(r.URL.Path)
	if match == nil {
		http.NotFound(w, r)
		return
	}
	title := match[2]

	//Do Save
	body := r.FormValue("body")
	err := v.repo.SaveWiki(&models.Wiki{Title: title, Body: []byte(body)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

//Edit handles the Edit action
func (v *WikiController) Edit(w http.ResponseWriter, r *http.Request) {
	//Check repo
	if v.repo == nil {
		http.Error(w, "Must initialize controller using MakeWikiController.", http.StatusInternalServerError)
		return
	}

	//Check path
	match := v.validPath.FindStringSubmatch(r.URL.Path)
	if match == nil {
		http.NotFound(w, r)
		return
	}
	title := match[2]

	//Do Edit
	wiki, err := v.repo.LoadWiki(title)
	if err != nil {
		//Wiki not found so create a new one fore editing
		wiki = &models.Wiki{Title: title}
	}

	wikiTitles, err := v.repo.ListWikiTitles()
	if v.repo != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	v.renderView(w, "edit", &models.Page{PageTitle: title, WikiTitles: wikiTitles, Wiki: wiki})
}
