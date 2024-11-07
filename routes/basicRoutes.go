package routes

import (
	"testproj/handlers"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitializeRoutes(userHandler *handlers.UserHandler, testsHandler *handlers.TestsHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/createuser", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/updateuser", userHandler.UpdatePassword).Methods("PUT")
	router.HandleFunc("/authorize", userHandler.AuthorizeUser).Methods("GET")
	router.HandleFunc("/testslist", testsHandler.GetListTests).Methods("GET")
	router.HandleFunc("/gettest", testsHandler.GetTest).Methods("GET")
	router.HandleFunc("/unpassedtest", testsHandler.GetUnpassedTests).Methods("GET")
	router.HandleFunc("/getmytests", testsHandler.GetMyTests).Methods("GET")
	router.HandleFunc("/addedtesttouser", testsHandler.AddedUserTest).Methods("POST")
	router.Handle("/metrics", promhttp.Handler()).Methods("GET") 
	return router
}
