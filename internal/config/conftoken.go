package config

import (
	"os"
)

type configToken struct {
	AccessTTL  string
	RefreshTTL string
	SecretKey  string
}

func TokenCFG() *configToken {

	return &configToken{
		AccessTTL:  os.Getenv("ACCESS_TTL"),
		RefreshTTL: os.Getenv("REFRESH_TOKEN"),
		SecretKey:  os.Getenv("SECRET_KEY"),
	}
}
