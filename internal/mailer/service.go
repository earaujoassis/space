package mailer

import (
	"github.com/earaujoassis/space/internal/config"
)

type Mailer struct {
	cfg *config.Config
}

func NewMailer(cfg *config.Config) *Mailer {
	return &Mailer{
		cfg: cfg,
	}
}
