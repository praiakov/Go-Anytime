package routes

import (
	"github.com/gorilla/mux"
	"github.com/praiakov/webapi/controllers"
)

func LoadRoutes(router *mux.Router) mux.Router {

	router.HandleFunc("/contato", controllers.GetPeople).Methods("GET")
	router.HandleFunc("/contato/{id}", controllers.GetPerson).Methods("GET")
	router.HandleFunc("/contato/{id}", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/contato/{id}", controllers.DeletePerson).Methods("DELETE")

	return *router

}
