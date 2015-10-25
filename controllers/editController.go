//Package controllers contains controllers for the web application.
package controllers

import (
	"net/http"

	"github.com/mikeyb214/higo/helpers"
	"github.com/mikeyb214/higo/models"
)

//EditController contains logic for the Edit action
func EditController(w http.ResponseWriter, r *http.Request, title string) {
	p, err := models.LoadPage(title)
	if err != nil {
		p = &models.Page{Title: title}
	}
	helpers.RenderView(w, "edit", p)
}
