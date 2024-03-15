package handlers

import (
	"encoding/base64"
	"net/http"
)

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("Refresh")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value

	tn, err := base64.StdEncoding.DecodeString(tknStr)
	if err != nil {
		h.logger.L.Error(err)
	}
	g, f := h.service.RefreshToken(string(tn))
	h.logger.L.WithField("handler.Refresh", "").Info("Значение g   ", g)
	if g {
		acces, refresh, err := h.service.GetTokens("", f)
		if err != nil {
			h.logger.L.WithField("handler.Refresh", "").Error(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		http.SetCookie(w, acces)
		http.SetCookie(w, refresh)

		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}
