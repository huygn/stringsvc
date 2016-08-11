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

	svc := stringsvc.NewStringService()
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

	svc := stringsvc.NewStringService()
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
