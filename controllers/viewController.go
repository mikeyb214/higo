//Package controllers contains controllers for the web application.
package controllers

import (
	"net/http"

	"github.com/mikeyb214/higo/helpers"
	"github.com/mikeyb214/higo/models"
)

//ViewController contains logic for the View action
func ViewController(w http.ResponseWriter, r *http.Request, title string) {
	p, err := models.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	helpers.RenderView(w, "view", p)
}
