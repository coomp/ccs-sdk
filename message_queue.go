package ccssdk

type MessageQueueRepository interface {
	Subscribe(topic string, funcs HandleFuncs)
}

type MessageQueueService struct {
	Repo MessageQueueRepository
}

func NewMessageQueueService(repo MessageQueueRepository) *MessageQueueService {
	return &MessageQueueService{
		Repo: repo,
	}
}

func (s *MessageQueueService) Subscribe(topic string, funcs HandleFuncs) {
	s.Repo.Subscribe(topic, funcs)
}
