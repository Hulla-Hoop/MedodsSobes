package service

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) GetTokens(reqId string, guid string) (*http.Cookie, *http.Cookie, error) {

	expirationTime := time.Now().Add(30 * time.Second)

	claims := &Claims{
		Username: users[guid],
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	jwt, err := token.SignedString([]byte("shamil"))
	if err != nil {
		s.logger.L.Error(err)
		return nil, nil, err
	}

	acces := &http.Cookie{
		Name:    "token",
		Value:   jwt,
		Expires: expirationTime,
	}

	timess := time.Now()

	hash, err := bcrypt.GenerateFromPassword([]byte(timess.String()), bcrypt.DefaultCost)
	if err != nil {
		return nil, nil, err
	}

	s.logger.L.WithField("Service.Gettoken", reqId).Info("sha3 string -----", string(hash))

	ref := base64.StdEncoding.EncodeToString(hash)

	s.logger.L.WithField("Service.Gettoken", reqId).Info("Base64 string -----", ref)

	refresh := &http.Cookie{
		Name:     "Refresh",
		Value:    ref,
		Expires:  time.Now().Add(time.Minute * 3),
		HttpOnly: true,
	}

	RT[string(hash)] = guid

	return acces, refresh, nil
}
