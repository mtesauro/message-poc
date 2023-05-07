package ports

import "message-poc/internal/core/domain"

type MessengerService interface {
	SaveMessage(message domain.Message) error
	ReadMessage(id string) (*domain.Message, error)
	ReadMessages() ([]*domain.Message, error)
	CloneMessage() (*domain.Message, error)
}

type MessengerRepository interface {
	SaveMessage(message domain.Message) error
	ReadMessage(id string) (*domain.Message, error)
	ReadMessages() ([]*domain.Message, error)
	//CloneMessage() (*domain.Message, error)
}
