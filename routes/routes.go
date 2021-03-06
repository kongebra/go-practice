package routes

import (
	"github.com/gorilla/mux"
	"github.com/kongebra/go-practice/db"
	"github.com/kongebra/go-practice/homepage"
	"log"
)

// Setup connects router, logger and database to controllers
func Setup(router *mux.Router, logger *log.Logger, db *db.Database) {
	h := homepage.New(logger, db)

	h.SetupRoutes(router)
}
