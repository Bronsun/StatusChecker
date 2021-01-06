package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Bronsun/StatusChecker/helpers"

	"github.com/Bronsun/StatusChecker/controllers"
	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()
	r.Use(helpers.PanicHandler)
	r.HandleFunc("/login", controllers.LoginHandler).Methods("POST")
	r.HandleFunc("/register", controllers.RegisterHandler).Methods("POST")
	r.HandleFunc("/user/{id}", controllers.GetUserHandler).Methods("GET")
	r.HandleFunc("/addLink", controllers.AddLinkHandler).Methods("POST")
	//r.HandleFunc("/status/{url}",controllers.StatusCheckerHandler).Methods("POST")
	fmt.Println("Server is up and running on port :4040")
	log.Fatal(http.ListenAndServe(":4040", r))
}
