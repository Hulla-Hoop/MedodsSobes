package handlers

import (
	"encoding/base64"
	"net/http"
)

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("reqID").(string)
	if !ok {
		reqID = ""
	}

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

	bcryptToken, err := base64.StdEncoding.DecodeString(tknStr)
	if err != nil {
		h.logger.L.WithField("Handler.Refresh", reqID).Error(err)
	}
	ok, guid := h.service.RefreshToken(reqID, string(bcryptToken))
	h.logger.L.WithField("Handler.Refresh", reqID).Info("Значение ok   ", ok)
	if ok {
		acces, refresh, err := h.service.GetTokens(reqID, guid)
		if err != nil {
			h.logger.L.WithField("handler.Refresh", reqID).Error(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		http.SetCookie(w, acces)
		http.SetCookie(w, refresh)

		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}
