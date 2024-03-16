package main

import (
	"medos/internal/DB/mongo"
	"medos/internal/handlers"
	"medos/internal/logger"
	"medos/internal/service"
	"net/http"
)

func main() {

	l := logger.New()
	m := mongo.New(l)
	s := service.New(l, m)
	h := handlers.New(l, s)

	http.HandleFunc("/rere", h.SignIn)
	http.HandleFunc("/test", h.Test)
	http.HandleFunc("/refresh", h.Refresh)

	http.ListenAndServe(":8080", nil)

}
