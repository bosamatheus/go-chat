package chat

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Ping() (string, error) {
	return s.repo.Pong()
}
