package handlers

import (
	"encoding/base64"
	"errors"
	"net/http"
)

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {

	reqID, oks := r.Context().Value("reqID").(string)
	if !oks {
		reqID = ""
	}

	guid := r.URL.Query().Get("guid")
	c, err := r.Cookie("Refresh")
	ok := errors.Is(err, http.ErrNoCookie)

	if ok {
		h.logger.L.WithField("Handler.SingIn", reqID).Debug("Полученый guid ---- ", guid)

		acces, refresh, err := h.service.GetTokens(reqID, guid)
		if err != nil {
			h.logger.L.WithField("Handler.SingIn", reqID).Error(err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		http.SetCookie(w, acces)
		http.SetCookie(w, refresh)
	} else {
		tknStr := c.Value

		bcryptToken, err := base64.StdEncoding.DecodeString(tknStr)
		if err != nil {
			h.logger.L.Error(err)
		}

		sess, err := h.service.ChekSess(reqID, string(bcryptToken))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		if sess.Guid == guid {
			w.Write([]byte("У вас есть рефреш токен перейдите по пути /refresh для обновления токена доступа "))

		} else {
			h.logger.L.Info(guid)

			acces, refresh, err := h.service.GetTokens(reqID, guid)
			if err != nil {
				h.logger.L.WithField("Handler.SingIn", reqID).Error(err)
				w.WriteHeader(http.StatusInternalServerError)
			}

			http.SetCookie(w, acces)
			http.SetCookie(w, refresh)
		}
	}
}
