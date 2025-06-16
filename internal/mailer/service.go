package mailer

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"

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

// SendEmail uses the AWS SES service to send e-mail messages
func (m *Mailer) SendEmail(subject, body, mailTo string) error {
	mailKey := strings.Split(m.cfg.MailerAccess, ":")
	mailFrom := m.cfg.MailFrom
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(mailKey[2]),
		Credentials: credentials.NewStaticCredentials(mailKey[0], mailKey[1], ""),
	})
	if err != nil {
		return err
	}

	svc := ses.New(sess)
	params := &ses.SendEmailInput{
		Source: aws.String(mailFrom),
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(mailTo),
			},
		},
		Message: &ses.Message{
			Subject: &ses.Content{
				Data:    aws.String(subject),
				Charset: aws.String("utf-8"),
			},
			Body: &ses.Body{
				Html: &ses.Content{
					Data:    aws.String(body),
					Charset: aws.String("utf-8"),
				},
			},
		},
		ReplyToAddresses: []*string{
			aws.String(mailFrom),
		},
	}
	if _, err := svc.SendEmail(params); err != nil {
		return err
	}

	return nil
}
