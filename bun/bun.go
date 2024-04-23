package bun

import (
	"context"
	"time"

	"github.com/uptrace/bun"

	"github.com/andiksetyawan/log"
)

// Field names
const (
	OperationFieldName     = "operation"
	OperationTimeFieldName = "operation_time_ms"
)

// QueryHook defines the
// structure of our query hook
// it implements the bun.QueryHook
// interface
type QueryHook struct {
	bun.QueryHook

	logger       log.Logger
	slowDuration time.Duration
}

// QueryHookOptions defines the
// available options for a new
// query hook.
type QueryHookOptions struct {
	Logger       log.Logger
	SlowDuration time.Duration
}

// NewQueryHook returns a new query hook for use with
// uptrace/bun.
func NewQueryHook(options QueryHookOptions) QueryHook {
	return QueryHook{
		logger:       options.Logger,
		slowDuration: options.SlowDuration,
	}
}

func (qh QueryHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

func (qh QueryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	queryDuration := time.Since(event.StartTime)

	attrs := []any{
		OperationFieldName, event.Operation(),
		OperationTimeFieldName, queryDuration.Milliseconds(),
	}

	// Errors will always be logged
	if event.Err != nil {
		attrs = append(attrs, event.Operation(), event.Err)
		qh.logger.Error(ctx, event.Query, attrs...)
		return
	}

	if queryDuration >= qh.slowDuration && qh.slowDuration != 0 {
		qh.logger.Info(ctx, event.Query, attrs...)
	} else {
		qh.logger.Debug(ctx, event.Query, attrs...)
	}
}
