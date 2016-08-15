package stringsvc

import (
	"time"

	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"
)

// Middleware describes a service middleware.
type Middleware func(Service) Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

func (mw loggingMiddleware) Log(method string, input, output interface{}, err error, took time.Duration) error {
	return mw.logger.Log(
		"method", method,
		"input", input,
		"output", output,
		"err", err,
		"took", took,
	)
}

func (mw loggingMiddleware) Uppercase(ctx context.Context, s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.Log(
			"uppercase",
			s,
			output,
			err,
			time.Since(begin),
		)
	}(time.Now())

	return mw.next.Uppercase(ctx, s)
}

func (mw loggingMiddleware) Count(ctx context.Context, s string) (n int) {
	defer func(begin time.Time) {
		mw.Log(
			"count",
			s,
			n,
			nil,
			time.Since(begin),
		)
	}(time.Now())

	return mw.next.Count(ctx, s)
}
