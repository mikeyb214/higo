//Package controllers contains controllers for the web application.
package controllers

import (
	"net/http"

	"github.com/mikeyb214/higo/models"
)

//SaveController contains logic for the Save action
func SaveController(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &models.Page{Title: title, Body: []byte(body)}
	err := p.SavePage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
