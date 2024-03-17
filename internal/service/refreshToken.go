package service

import "medods/internal/model"

// возращает guid пользователя если его сессия есть в базе
func (s *Service) RefreshToken(reqID string, token string) (bool, string) {
	s.logger.L.WithField("service.RefreshToken", reqID).Debug(token)
	session, err := s.ChekSess(reqID, token)
	if err != nil {
		return false, ""
	} else {
		s.db.DeleteSess(reqID, token)
		s.logger.L.WithField("service.RefreshToken", reqID).Debug(session)
		return true, session.Guid
	}
}

// проверка наличия сессии в монго
func (s *Service) ChekSess(reqId string, token string) (*model.Session, error) {
	session, err := s.db.ChekSess(reqId, token)
	return session, err
}
