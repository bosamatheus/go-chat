package chat

type Repository interface {
	KeyNotExists(string) bool
	GetPreviousValues(string) ([]string, error)
	SaveValue(string, []byte) error
}

type UseCase interface {
	ChatExists(string) bool
	GetChatHistory(string) ([]string, error)
	SaveMessage(string, []byte) error
}
