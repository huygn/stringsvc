package stringsvc_test

import (
	"bytes"
	"flag"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"

	"github.com/gnhuy91/stringsvc"
)

var debug bool

func init() {
	flag.BoolVar(&debug, "debug", false, "")
	flag.Parse()
}

type service struct {
	UppercaseF func(context.Context, string) (string, error)
	CountF     func(context.Context, string) int
}

func (svc *service) Uppercase(ctx context.Context, s string) (string, error) {
	return svc.UppercaseF(ctx, s)
}

func (svc *service) Count(ctx context.Context, s string) (n int) {
	return svc.CountF(ctx, s)
}

var (
	logfmtRegex = `^method=(\S+) ` +
		`input=("*[a-zA-Z0-9_ ]*"*) ` +
		`output=("*[a-zA-Z0-9_ ]*"*) ` +
		`err=("*[a-zA-Z0-9_ ]*"*) ` +
		`took=(\S+)$`
	methodPrefix = `method=`
)

func TestLogMiddlewareUppercase(t *testing.T) {
	ctx := context.Background()

	var buf bytes.Buffer
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(&buf)
		if debug {
			logger = log.NewLogfmtLogger(io.MultiWriter(&buf, os.Stderr))
		}
	}

	svc := &service{
		UppercaseF: func(ctx context.Context, s string) (string, error) {
			return "", nil
		},
	}
	logMw := stringsvc.LoggingMiddleware(logger)(svc)
	logMw.Uppercase(ctx, "")

	testLogFmt(t, buf, "uppercase")
}

func TestLogMiddlewareCount(t *testing.T) {
	ctx := context.Background()

	var buf bytes.Buffer
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(&buf)
		if debug {
			logger = log.NewLogfmtLogger(io.MultiWriter(&buf, os.Stderr))
		}
	}

	svc := &service{
		CountF: func(ctx context.Context, s string) int {
			return 0
		},
	}
	logMw := stringsvc.LoggingMiddleware(logger)(svc)
	logMw.Count(ctx, "")

	testLogFmt(t, buf, "count")
}

func testLogFmt(t *testing.T, logOutput bytes.Buffer, method string) {
	b, err := ioutil.ReadAll(&logOutput)
	if err != nil {
		t.Error(err)
	}
	b = bytes.TrimSpace(b)
	match, err := regexp.Match(logfmtRegex, b)
	if err != nil {
		t.Error(err)
	}
	if !match {
		t.Errorf("log output does not match regex %q, ouput %q", logfmtRegex, b)
	}
	prefix := []byte(methodPrefix + method)
	if !bytes.HasPrefix(b, prefix) {
		t.Errorf("log output method does not match, want prefix %q, got %q", prefix, b)
	}
}
