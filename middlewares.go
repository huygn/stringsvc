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
	logfmt := logfmtStruct{
		method,
		input,
		output,
		err,
		took,
	}
	return mw.logger.Log(logfmt.keyvals()...)
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

// logfmtStruct contains logfmt-style logging fields
type logfmtStruct struct {
	method        string
	input, output interface{}
	err           error
	took          time.Duration
}

// keyvals return key-val pairs to feed to go-kit/log.Logger.Log() method
func (s *logfmtStruct) keyvals() []interface{} {
	return []interface{}{
		"method", s.method,
		"input", s.input,
		"output", s.output,
		"err", s.err,
		"took", s.took,
	}
}
