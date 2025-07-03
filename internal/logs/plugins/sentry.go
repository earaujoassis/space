package plugins

import (
	"crypto/tls"
	"errors"
	"net/http"

	"github.com/getsentry/sentry-go"

	"github.com/earaujoassis/space/internal/logs/level"
)

var isSentryAvailable bool

func init() {
	isSentryAvailable = false
}

func SetupSentry(environment, release, sentryUrl string) error {
	isSentryAvailable = false
	if sentryUrl != "" {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		err := sentry.Init(sentry.ClientOptions{
			Dsn:              sentryUrl,
			TracesSampleRate: 1.0,
			Environment:      environment,
			Release:          release,
			HTTPClient:       client,
		})
		if err != nil {
			return err
		} else {
			isSentryAvailable = true
		}
	}

	return nil
}

func IsSentryAvailable() bool {
	return isSentryAvailable
}

func CaptureMessage(msg string, lev level.Level) {
	if isSentryAvailable {
		sentry.WithScope(func(scope *sentry.Scope) {
			switch lev {
			case level.Critical, level.Panic:
				scope.SetLevel(sentry.LevelFatal)
			case level.Error:
				scope.SetLevel(sentry.LevelError)
			default:
				scope.SetLevel(sentry.LevelInfo)
			}
			sentry.CaptureMessage(msg)
		})
	}
}

func CaptureException(msg string) {
	if isSentryAvailable {
		err := errors.New(msg)
		sentry.CaptureException(err)
	}
}
