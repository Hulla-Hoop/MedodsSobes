package service

import "medos/internal/model"

func (s *Service) RefreshToken(token string) (bool, string) {
	s.logger.L.WithField("service.RefreshToken", "").Info(token)
	session, err := s.ChekSess("", token)
	if err != nil {
		return false, ""
	} else {
		s.db.DeleteSess("", token)
		s.logger.L.Info(session)
		return true, session.Guid
	}
}

func (s *Service) ChekSess(reqId string, token string) (*model.Session, error) {
	session, err := s.db.ChekSess("", token)
	return session, err
}
