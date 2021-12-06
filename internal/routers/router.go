package routers

import (
	"github.com/Protoc0l7/simple_server/internal/controllers"
	"github.com/gorilla/mux"
)

func InitRoutes(controller *controllers.Controller) *mux.Router {
	router := mux.NewRouter()
	router.Path("/user").HandlerFunc(controller.CreateUser).Methods("POST")
	router.Path("/user/{userID}").HandlerFunc(controller.GetAllUserMessages).Methods("Get")
	router.Path("/message").HandlerFunc(controller.CreateMessage).Methods("POST")
	router.Path("/message/{messageID}").HandlerFunc(controller.GetMessage).Methods("GET")
	return router
}