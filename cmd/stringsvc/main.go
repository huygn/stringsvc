package main

import (
	"golang.org/x/net/context"

	"github.com/gnhuy91/stringsvc"
)

func main() {
	ctx := context.Background()
	var svc stringsvc.Service

	stringsvc.MakeHTTPHandler(ctx, svc)
	// log.Fatal(http.ListenAndServe(":8080", h))

	// var logger log.Logger
	// {
	// 	logger = log.NewLogfmtLogger(os.Stderr)
	// 	logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
	// 	logger = log.NewContext(logger).With("caller", log.DefaultCaller)
	// }

	// var ctx context.Context
	// {
	// 	ctx = context.Background()
	// }

	// var s stringsvc.Service

	// var h http.Handler
	// {
	// 	h = stringsvc.MakeHTTPHandler(ctx, s)
	// }

	// errs := make(chan error)
	// go func() {
	// 	c := make(chan os.Signal)
	// 	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	// 	errs <- fmt.Errorf("%s", <-c)
	// }()

	// go func() {
	// 	httpAddr := ":8080"
	// 	logger.Log("transport", "HTTP", "addr", httpAddr)
	// 	errs <- http.ListenAndServe(httpAddr, h)
	// }()

	// logger.Log("exit", <-errs)
}
