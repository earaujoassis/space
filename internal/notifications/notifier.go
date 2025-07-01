package notifications

import (
	"encoding/json"

	"github.com/hibiken/asynq"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/logs"
	"github.com/earaujoassis/space/internal/utils"
	"github.com/earaujoassis/space/internal/workers"
)

type Notifier struct {
	cfg *config.Config
}

func NewNotifier(cfg *config.Config) *Notifier {
	return &Notifier{
		cfg: cfg,
	}
}

// Announce is used to communicate actions throughout the application,
//
//	using e-mail messages (production-only) or stdout (development-only)
func (n *Notifier) Announce(name string, data utils.H) {
	switch n.cfg.Environment {
	case config.Production:
		cfg := n.cfg
		enqueuer := asynq.NewClient(asynq.RedisClientOpt{
			Addr: cfg.MemoryDNS(),
			DB:   cfg.MemorystoreIndex,
		})
		defer enqueuer.Close()
		payload, err := json.Marshal(workers.EmailDeliveryPayload{
			Name: name,
			Data: data,
		})
		if err != nil {
			logs.Propagatef(logs.Error, "could not enqueue task for email delivery: %s", name)
			return
		}
		enqueuer.Enqueue(asynq.NewTask(workers.TypeEmailDelivery, payload))
	default:
		logs.Propagatef(logs.Info, "Action `%s` with data `%v`\n", name, data)
	}
}
