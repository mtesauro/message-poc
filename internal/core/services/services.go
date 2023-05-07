package services

import (
	"message-poc/internal/core/domain"
	"message-poc/internal/core/ports"

	"github.com/google/uuid"
)

type MessengerService struct {
	repo ports.MessengerRepository
}

func NewMessengerService(repo ports.MessengerRepository) *MessengerService {
	return &MessengerService{
		repo: repo,
	}
}

func (m *MessengerService) SaveMessage(message domain.Message) error {
	message.ID = uuid.New().String()
	return m.repo.SaveMessage(message)
}

func (m *MessengerService) ReadMessage(id string) (*domain.Message, error) {
	return m.repo.ReadMessage(id)
}

func (m *MessengerService) ReadMessages() ([]*domain.Message, error) {
	return m.repo.ReadMessages()
}

func (m *MessengerService) CloneMessage(origMess domain.Message) error {
	message := domain.Message{
		ID:     uuid.New().String(),
		Body:   origMess.Body,
		Status: origMess.Status,
	}
	return m.repo.SaveMessage(message)
}
