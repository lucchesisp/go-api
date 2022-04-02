package routes

import (
	"github.com/gorilla/mux"
	"github.com/lucchesisp/go-api/controllers"
	"github.com/lucchesisp/go-api/middleware"
	"log"
	"net/http"
)

func HandleRequest() {

	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/personalidades/{id}", controllers.DeletandoUmaPersonalidade).Methods("Delete")
	r.HandleFunc("/personalidades/{id}", controllers.EditaPersonalidade).Methods("Put")

	log.Fatal(http.ListenAndServe(":3000", r))
}
