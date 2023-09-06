package plugins

import (
	"crypto/tls"
	"errors"
	"net/http"

	"github.com/getsentry/sentry-go"
)

var isSentryAvailable bool

func init() {
	isSentryAvailable = false
}

func SetupSentry(environment, sentryUrl string) error {
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

func CaptureException(msg string) {
	if isSentryAvailable {
		sentry.CaptureException(errors.New(msg))
	}
}
