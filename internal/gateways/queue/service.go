package queue

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/hibiken/asynq"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/logs"
)

type QueueService struct {
	config   *config.Config
	client   *asynq.Client
	provider *miniredis.Miniredis
}

func NewQueueService(cfg *config.Config) *QueueService {
	var addr string
	var provider *miniredis.Miniredis
	var err error

	switch cfg.Environment {
	case config.Production, config.Development, config.Integration:
		addr = cfg.MemoryDNS()
		provider = nil
	case config.Test:
		provider, err = miniredis.Run()
		if err != nil {
			logs.Propagate(logs.LevelPanic, err.Error())
		}
		addr = provider.Addr()
	default:
		logs.Propagate(logs.LevelPanic, "gateway misconfigured")
	}

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
	if qs.provider != nil {
		qs.provider.Close()
	}
}
