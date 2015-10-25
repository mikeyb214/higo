//Package implements a simple wiki HTTP server.
package main

import (
	"net/http"

	"github.com/mikeyb214/higo/controllers"
	"github.com/mikeyb214/higo/helpers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/index.html")
	})
	http.HandleFunc("/view/", helpers.MakeController(controllers.ViewController))
	http.HandleFunc("/edit/", helpers.MakeController(controllers.EditController))
	http.HandleFunc("/save/", helpers.MakeController(controllers.SaveController))

	http.ListenAndServe(":8080", nil)
}
