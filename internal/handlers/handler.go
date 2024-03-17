package handlers

import (
	"encoding/base64"
	"fmt"
	"medods/internal/config"
	"medods/internal/logger"
	"medods/internal/service"
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

// Функция для тестирования: берет токен из куки и распечатывает имя пользователя из пэйлоуда
// также логирует рефреш токен
func (h *Handler) Test(w http.ResponseWriter, r *http.Request) {

	reqID := r.Context().Value("reqID").(string)
	if reqID == "" {
		reqID = ""
	}

	//проверка наличиия токена доступа
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

	cfg := config.TokenCFG()

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (any, error) {
		return []byte(cfg.SecretKey), nil
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
	//проверка наличия рефреш токена
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

	h.logger.L.WithField("Handler.Test", reqID).Debug("Sha3 strinng from test --", string(re))

	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}
