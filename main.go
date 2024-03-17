package main

import (
	"medos/internal/DB/mongo"
	"medos/internal/handlers"
	"medos/internal/logger"
	"medos/internal/service"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {

	l := logger.New()

	err := godotenv.Load()
	if err != nil {
		l.L.Info("Не загружается .env файл")
	}

	m := mongo.New(l)
	s := service.New(l, m)
	h := handlers.New(l, s)

	go s.ClearSession()

	http.HandleFunc("/signin", h.ReqID(h.SignIn))
	http.HandleFunc("/test", h.ReqID(h.Test))
	http.HandleFunc("/refresh", h.ReqID(h.Refresh))

	http.ListenAndServe(":8080", nil)

}
