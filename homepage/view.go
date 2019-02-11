package homepage

import (
	"github.com/kongebra/go-practice/models"
	"html/template"
	"net/http"
)

// PageData struct
type PageData struct {
	StatusCode int           `json:"statusCode"`
	PageTitle  string        `json:"pageTitle"`
	Posts      []models.Post `json:"posts"`
	tmpl       string
}

func (ctrl *Controller) view(w http.ResponseWriter, r *http.Request, data PageData) {
	w.WriteHeader(data.StatusCode)

	temp, err := template.ParseFiles("templates/layout.html", "templates/navbar.html", data.tmpl)

	if err != nil {
		ctrl.logger.Printf("could not parse files: %v\n", err)
	}

	if err = temp.ExecuteTemplate(w, "layout", data); err != nil {
		ctrl.logger.Printf("could not execute template: %v\n", err)
	}
}
