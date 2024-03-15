package service

var RT = map[string]string{}

func (s *Service) RefreshToken(token string) (bool, string) {
	s.logger.L.WithField("service.RefreshToken", "").Info(RT)
	rt, ok := RT[token]
	if !ok {
		return false, ""
	} else {
		delete(RT, token)
		s.logger.L.Info(rt)
		return true, rt
	}
}
