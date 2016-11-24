package mailer

import (
    "strings"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ses"

    "github.com/earaujoassis/space/config"
)

func SendEmail(subject, body, mail_to string) error {
    mail_key := strings.Split(config.GetConfig("SPACE_MAIL_ACCESS"), ":")
    mail_from := config.GetConfig("SPACE_MAIL_FROM")
    sess, err := session.NewSession(&aws.Config{
        Region:      aws.String(mail_key[2]),
        Credentials: credentials.NewStaticCredentials(mail_key[0], mail_key[1], ""),
    })
    if err != nil {
        return err
    }

    svc := ses.New(sess)
    params := &ses.SendEmailInput{
        Source: aws.String(mail_from),
        Destination: &ses.Destination{
            ToAddresses: []*string{
                aws.String(mail_to),
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
            aws.String(mail_from),
        },
    }
    if _, err := svc.SendEmail(params); err != nil {
        return err
    }

    return nil
}
