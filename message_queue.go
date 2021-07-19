package ccssdk

type MessageQueueRepository interface {
	SubscribeRequest(topic string, funcs RequestHandleFuncs) error
	SubscribeResponse(topic string, funcs ResponseHandleFuncs) error
	Start() error
}

type MessageQueueService struct {
	Repo MessageQueueRepository
}

func NewMessageQueueService(repo MessageQueueRepository) *MessageQueueService {
	return &MessageQueueService{
		Repo: repo,
	}
}

func (s *MessageQueueService) SubscribeRequest(topic string, funcs RequestHandleFuncs) error {
	return s.Repo.SubscribeRequest(topic, funcs)
}

func (s *MessageQueueService) SubscribeResponse(topic string, funcs ResponseHandleFuncs) error {
	return s.Repo.SubscribeResponse(topic, funcs)
}

func (s *MessageQueueService) Start() error {
	return s.Repo.Start()
}
