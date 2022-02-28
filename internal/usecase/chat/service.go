package chat

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) ChatExists(key string) bool {
	return !s.repo.KeyNotExists(key)
}

func (s *Service) GetChatHistory(key string) ([]string, error) {
	return s.repo.GetPreviousValues(key)
}

func (s *Service) SaveMessage(key string, val []byte) error {
	return s.repo.SaveValue(key, val)
}
