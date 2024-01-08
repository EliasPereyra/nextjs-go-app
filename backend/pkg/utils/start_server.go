package utils

import (
	"api/pkg/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer(router *mux.Router) {
	routerWithMiddleware := middleware.EnableCORS(middleware.JsonContentTypeMiddleware(router))

  log.Fatal(http.ListenAndServe(":8000", routerWithMiddleware))
}