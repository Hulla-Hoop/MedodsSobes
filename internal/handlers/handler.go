package handlers

import (
	"encoding/base64"
	"fmt"
	"medos/internal/logger"
	"medos/internal/service"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

type Handler struct {
	logger  *logger.Logger
	service *service.Service
}

func New(log *logger.Logger, serv *service.Service) *Handler {
	return &Handler{
		logger:  log,
		service: serv,
	}
}

func (h *Handler) Test(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {

			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tknStr := c.Value

	claims := &service.Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (any, error) {
		return []byte("shamil"), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	ref, err := r.Cookie("Refresh")
	if err != nil {
		if err == http.ErrNoCookie {

			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	re, err := base64.StdEncoding.DecodeString(ref.Value)
	if err != nil {
		h.logger.L.Error(err)
	}

	h.logger.L.Info("Sha3 strinng from test --", string(re))

	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}
