package service

import (
	"encoding/base64"
	"medos/internal/config"
	"medos/internal/model"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) GetTokens(reqId string, guid string) (*http.Cookie, *http.Cookie, error) {

	cfg := config.TokenCFG()

	acces, err := s.createAccessToken(cfg.AccessTTL, cfg.SecretKey, guid)
	if err != nil {
		return nil, nil, err
	}

	timess := time.Now()

	hash, refresh, err := s.createRefreshToken(cfg.RefreshTTL, timess.String())
	if err != nil {
		return nil, nil, err
	}

	var session model.Session

	session.TimeCreatedTocken = timess.String()
	session.BcryptTocken = string(hash)
	session.Guid = guid

	s.db.CreateSess("", &session)

	return acces, refresh, nil
}

func (s *Service) createAccessToken(accessTTL string, secret string, guid string) (*http.Cookie, error) {
	TTL, err := strconv.Atoi(accessTTL)
	if err != nil {
		return nil, err
	}

	expirationTime := time.Now().Add(time.Minute * time.Duration(TTL))

	claims := &Claims{
		Username: users[guid],
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	jwt, err := token.SignedString([]byte(secret))
	if err != nil {
		s.logger.L.Error(err)
		return nil, err
	}

	acces := &http.Cookie{
		Name:    "token",
		Value:   jwt,
		Expires: expirationTime,
	}
	return acces, nil
}

func (s *Service) createRefreshToken(refreshTTL string, times string) (string, *http.Cookie, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(times), bcrypt.DefaultCost)
	if err != nil {
		return "", nil, err
	}

	s.logger.L.WithField("Service.Gettoken", "").Info("sha3 string -----", string(hash))

	ref := base64.StdEncoding.EncodeToString(hash)

	s.logger.L.WithField("Service.Gettoken", "").Info("Base64 string -----", ref)

	TTL, err := strconv.Atoi(refreshTTL)
	if err != nil {
		return "", nil, err
	}

	refresh := &http.Cookie{
		Name:     "Refresh",
		Value:    ref,
		Expires:  time.Now().Add(time.Minute * time.Duration(TTL)),
		HttpOnly: true,
	}
	return string(hash), refresh, nil
}
