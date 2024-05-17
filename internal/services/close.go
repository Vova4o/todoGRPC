package services

func (s *Service) Close() string {
	s.DB.CloseDB()
	return "Соединение с базой данных закрыто."
}
