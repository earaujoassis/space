package workers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/mailer"
	"github.com/earaujoassis/space/internal/utils"
)

const (
	TypeEmailDelivery = "email:delivery"
)

type EmailDeliveryPayload struct {
	Name string
	Data utils.H
}

type EmailDeliveryProcessor struct {
	mailerService *mailer.Mailer
}

func NewEmailDeliveryProcessor(cfg *config.Config) *EmailDeliveryProcessor {
	ms := mailer.NewMailer(cfg)
	return &EmailDeliveryProcessor{
		mailerService: ms,
	}
}

func (p *EmailDeliveryProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var payload EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("payload deserealization failed: %v: %w", err, asynq.SkipRetry)
	}
	return p.mailerService.AnnounceMessage(payload.Name, payload.Data)
}
