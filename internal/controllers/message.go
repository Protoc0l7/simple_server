package controllers

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"net/http"
)

type Controller struct{
	DB *redis.Client
}

func (c Controller) CreateMessage(w http.ResponseWriter, r *http.Request)  {
	result, _ := c.DB.Ping().Result()
	fmt.Println(result)
}

func (c Controller) GetMessage(w http.ResponseWriter, r *http.Request)  {
	messageID := mux.Vars(r)["messageID"]
	fmt.Println("Hello")
}