package main

import (
	"github.com/Protoc0l7/simple_server/internal/controllers"
	"github.com/Protoc0l7/simple_server/internal/db"
	"github.com/Protoc0l7/simple_server/internal/routers"
	"log"
	"net/http"
)

func main() {
	db := db.NewClient()
	controller := controllers.Controller{
		DB: db,
	}

	router := routers.InitRoutes(&controller)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}