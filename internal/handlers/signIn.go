package handlers

import (
	"net/http"
)

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {

	str := r.URL.Query().Get("guid")
	h.logger.L.Info(str)

	acces, refresh, err := h.service.GetTokens("", str)
	if err != nil {
		h.logger.L.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	http.SetCookie(w, acces)
	http.SetCookie(w, refresh)
}
