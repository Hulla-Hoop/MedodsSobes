package service

import "time"

// Чистит базу от устаревших сессий
func (s *Service) ClearSession() {
	for {
		time.Sleep(time.Minute * 5)
		s.db.DeleteOld()
	}
}
