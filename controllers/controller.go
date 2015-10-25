//Package controllers contains controllers for the web application.
package controllers

import (
	"html/template"
	"net/http"

	"github.com/mikeyb214/higo/models"
)

//Controller contains base propropties and methods for Controllers.
type Controller struct {
	repo     models.WikiRepository
	template *template.Template
}

//RenderView builds the HTML response from a template.
func (v *Controller) renderView(w http.ResponseWriter, view string, p *models.Page) {
	err := v.template.ExecuteTemplate(w, view+".gohtml", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
