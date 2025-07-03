package notifications

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/gateways/queue"
	"github.com/earaujoassis/space/internal/logs"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/repository"
	"github.com/earaujoassis/space/internal/utils"
	"github.com/earaujoassis/space/internal/workers"
)

type Notifier struct {
	cfg          *config.Config
	repositories *repository.RepositoryManager
}

func NewNotifier(cfg *config.Config, repositories *repository.RepositoryManager) *Notifier {
	return &Notifier{
		cfg:          cfg,
		repositories: repositories,
	}
}

// Announce is used to communicate actions throughout the application,
//
//	using e-mail messages (production-only) or stdout (development-only)
func (n *Notifier) Announce(user models.User, name string, data utils.H) {
	if !n.shouldNotify(name, user) {
		return
	}

	cfg := n.cfg
	switch cfg.Environment {
	case config.Production, config.Test:
		enqueuer := queue.NewQueueService(cfg)
		defer enqueuer.Close()
		payload, err := json.Marshal(workers.EmailDeliveryPayload{
			Name: name,
			Data: data,
		})
		if err != nil {
			logs.Propagatef(logs.LevelError, "could not enqueue task for email delivery: %s", name)
			return
		}
		info, err := enqueuer.Enqueue(workers.TypeEmailDelivery, payload)
		fmt.Printf("%v\n", err)
		fmt.Printf("%v\n", info)
	default:
		logs.Propagatef(logs.LevelInfo, "Action `%s` with data `%v`\n", name, data)
	}
}

func (n *Notifier) shouldNotify(name string, user models.User) bool {
	notificationSettingsKey := mapNotificationNameToSettings(name)
	if notificationSettingsKey == "" {
		return true
	}

	parts := strings.Split(notificationSettingsKey, ".")
	if len(parts) != 3 {
		return true
	}

	repositories := n.repositories
	setting := repositories.Settings().FindOrDefault(user, parts[0], parts[1], parts[2])
	value, _ := setting.DeserializeValue()
	return value.(bool)
}
