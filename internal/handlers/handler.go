package handlers

import "medos/internal/logger"

type Handler struct {
	logger logger.Logger
}

func NewHanler(log logger.Logger) *Handler {
	return &Handler{
		logger: log,
	}
}
