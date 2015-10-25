//Package implements a simple wiki web application.
package main

import (
	"html/template"
	"net/http"

	"github.com/mikeyb214/higo/controllers"
	"github.com/mikeyb214/higo/models"
)

var viewTemplate = template.Must(template.ParseFiles("./views/index.gohtml", "./views/edit.gohtml", "./views/view.gohtml"))

func main() {
	repo := new(models.WikiFileRepository)
	homeController := controllers.MakeHomeController(repo, viewTemplate)
	wikiController := controllers.MakeWikiController(repo, viewTemplate)

	http.HandleFunc("/", homeController.Render)
	http.HandleFunc("/view/", wikiController.Render)
	http.HandleFunc("/edit/", wikiController.Edit)
	http.HandleFunc("/save/", wikiController.Save)

	http.ListenAndServe(":8080", nil)
}
