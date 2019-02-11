package main

import (
	"github.com/gorilla/mux"
	"github.com/kongebra/go-practice/db"
	"github.com/kongebra/go-practice/routes"
	"github.com/kongebra/go-practice/server"
	"github.com/kongebra/go-practice/utils"
	"log"
	"net/http"
	"os"
)

func main() {
	// Create logger
	logger := log.New(os.Stdout, "main: ", log.LstdFlags|log.Lshortfile)

	DB := db.New("root:@tcp(localhost:3306)/bsc")

	// Create router
	router := mux.NewRouter().StrictSlash(true)

	// Setup router
	routes.Setup(router, logger, DB)

	// Set path prefix for the assets-folder
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	// Create server
	srv := server.New(router, utils.GetPort())

	// Logging startup info
	logger.Printf("starting up server, listening on port%s\n", utils.GetPort())

	// Startup server
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("server fail to start: %v\n", err)
	}
}
