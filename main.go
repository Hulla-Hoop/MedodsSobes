package main

import (
	"medos/internal/handlers"
	"medos/internal/logger"
	"medos/internal/service"
	"net/http"
)

func main() {
	l := logger.New()
	s := service.New(l)
	h := handlers.New(l, s)

	http.HandleFunc("/rere", h.SignIn)
	http.HandleFunc("/test", h.Test)
	http.HandleFunc("/refresh", h.Refresh)

	http.ListenAndServe(":8080", nil)

}
