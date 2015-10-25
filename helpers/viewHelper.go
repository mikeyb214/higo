//Package helpers contains helper methods that support other packages.
package helpers

import (
	"net/http"
	"text/template"

	"github.com/mikeyb214/higo/models"
)

var views = template.Must(template.ParseFiles("./views/edit.gohtml", "./views/view.gohtml"))

//RenderView builds the HTML response from a template.
func RenderView(w http.ResponseWriter, tmpl string, p *models.Page) {
	err := views.ExecuteTemplate(w, tmpl+".gohtml", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
