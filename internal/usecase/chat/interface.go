package chat

type Repository interface {
	Pong() (string, error)
}

type UseCase interface {
	Ping() (string, error)
}
