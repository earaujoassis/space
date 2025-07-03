package queue

import (
	"github.com/hibiken/asynq"

	"github.com/earaujoassis/space/internal/config"
)

type QueueService struct {
	config *config.Config
	client *asynq.Client
}

func NewQueueService(cfg *config.Config) *QueueService {
	addr := cfg.MemoryDNS()
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr: addr,
		DB:   cfg.MemorystoreIndex,
	})
	return &QueueService{
		config: cfg,
		client: client,
	}
}

func (qs *QueueService) Enqueue(typename string, payload []byte) (*asynq.TaskInfo, error) {
	return qs.client.Enqueue(asynq.NewTask(typename, payload))
}

func (qs *QueueService) Close() {
	qs.client.Close()
}
