package services

import "github.com/coomp/ccs-sdk/handle"

type MessageQueueRepository interface {
	Subscribe(topic string, funcs handle.HandleFuncs)
}

type MessageQueueService struct {
	Repo MessageQueueRepository
}

func NewMessageQueueService(repo MessageQueueRepository) *MessageQueueService {
	return &MessageQueueService{
		Repo: repo,
	}
}

func (s *MessageQueueService) Subscribe(topic string, funcs handle.HandleFuncs) {
	s.Repo.Subscribe(topic, funcs)
}
