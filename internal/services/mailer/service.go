package mailer

import (
    "strings"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ses"

    "github.com/earaujoassis/space/internal/config"
)

// SendEmail uses the AWS SES service to send e-mail messages
func SendEmail(subject, body, mailTo string) error {
    var cfg config.Config = config.GetGlobalConfig()
    mailKey := strings.Split(cfg.MailerAccess, ":")
    mailFrom := cfg.MailFrom
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
