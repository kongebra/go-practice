package homepage

import (
	"github.com/gorilla/mux"
	"github.com/kongebra/go-practice/db"
	"log"
	"net/http"
	"strconv"
)

const (
	aboutMessage = "about Us"
)

// Controller struct
type Controller struct {
	logger *log.Logger
	db     *db.Database
}

// index
func (ctrl *Controller) index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8") // OP

	posts, err := ctrl.db.GetAllPosts()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	ctrl.view(w, r, PageData{
		StatusCode: http.StatusOK,
		PageTitle:  posts[0].Title,
		Posts:      posts,
		tmpl:       "templates/home.html",
	})
}

// single
func (ctrl *Controller) single(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8") // OP

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	posts, err := ctrl.db.GetSinglePost(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	ctrl.view(w, r, PageData{
		StatusCode: http.StatusOK,
		PageTitle:  "Frontpage",
		Posts:      posts,
		tmpl:       "templates/single.html",
	})
}

// about
func (ctrl *Controller) about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8") // OP
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(aboutMessage))
}

// SetupRoutes ...
func (ctrl *Controller) SetupRoutes(r *mux.Router) {
	r.HandleFunc("/", ctrl.index)
	r.HandleFunc("/post/{id:[0-9]+}", ctrl.single)
	r.HandleFunc("/about", ctrl.about)
}

// New creates a new Controller and returns a pointer
func New(logger *log.Logger, db *db.Database) *Controller {
	return &Controller{
		logger: logger,
		db:     db,
	}
}
