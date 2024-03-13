package handlers

import (
	"encoding/json"
	"net/http"
)

var users = map[string]string{
	"user1": "3825c945-8843-4b7d-995e-30b16c173c65",
	"user2": "019ed7ca-8286-40b8-ac80-1950c92dccfd",
}

type Credentials struct {
	GUID string
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
