package plugins

import (
	"context"
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/hibiken/asynq"
)

func SentryErrorHandler() asynq.ErrorHandlerFunc {
	return asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
		hub := sentry.CurrentHub().Clone()
		hub.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetTag("task_type", task.Type())
			scope.SetTag("error_handler", "fallback")
			scope.SetLevel(sentry.LevelError)
		})
		hub.CaptureMessage(fmt.Sprintf("[space-worker] Error: %s", err))
	})
}

func SentryMiddleware() asynq.MiddlewareFunc {
	return func(next asynq.Handler) asynq.Handler {
		return asynq.HandlerFunc(func(ctx context.Context, task *asynq.Task) error {
			hub := sentry.CurrentHub().Clone()

			hub.ConfigureScope(func(scope *sentry.Scope) {
				scope.SetTag("task_type", task.Type())
				scope.SetTag("component", "asynq_worker")
				scope.SetContext("task_payload", map[string]interface{}{
					"type":    task.Type(),
					"payload": string(task.Payload()),
				})
			})

			defer func() {
				if err := recover(); err != nil {
					hub.CaptureException(fmt.Errorf("panic in task %s: %v", task.Type(), err))
					hub.Flush(2 * time.Second)
					panic(err)
				}
			}()

			err := next.ProcessTask(ctx, task)

			if err != nil {
				hub.ConfigureScope(func(scope *sentry.Scope) {
					scope.SetLevel(sentry.LevelError)
					scope.SetContext("error_details", map[string]interface{}{
						"error":     err.Error(),
						"task_type": task.Type(),
					})
				})
				hub.CaptureException(err)
				hub.Flush(2 * time.Second)
			}

			return err
		})
	}
}
