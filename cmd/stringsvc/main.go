package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"

	"github.com/gnhuy91/stringsvc"
)

func main() {
	ctx := context.Background()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
		logger = log.NewContext(logger).With("caller", log.DefaultCaller)
	}

	var svc stringsvc.Service
	{
		svc = stringsvc.NewStringService()
		svc = stringsvc.LoggingMiddleware(logger)(svc)
	}

	h := stringsvc.MakeHTTPHandler(ctx, svc)
	logger.Log("exit", http.ListenAndServe(":8080", h))
}
