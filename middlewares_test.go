package stringsvc_test

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"

	"github.com/gnhuy91/stringsvc"
)

// For showing log output to Stderr during test.
// Usage:
//    go test -v -debug
//
var debug bool

func init() {
	flag.BoolVar(&debug, "debug", false, "")
	flag.Parse()
}

func TestLogMiddlewareUppercase(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{
			in:  `hello, world`,
			out: `method=uppercase input="hello, world" output="HELLO, WORLD" err=null`,
		},
		{
			in:  ``,
			out: `method=uppercase input= output= err="Empty string"`,
		},
	}

	svc := stringsvc.NewStringService()
	ctx := context.Background()

	for _, c := range cases {
		t.Run(fmt.Sprintf("in=%q", c.in), func(t *testing.T) {
			var buf bytes.Buffer
			var logger log.Logger
			{
				logger = log.NewLogfmtLogger(&buf)
				if debug {
					logger = log.NewLogfmtLogger(io.MultiWriter(&buf, os.Stderr))
				}
			}

			logMw := stringsvc.LoggingMiddleware(logger)(svc)
			logMw.Uppercase(ctx, c.in)
			testLogFmt(t, buf, c.out)
		})
	}
}

func TestLogMiddlewareCount(t *testing.T) {
	c := struct {
		in, out string
	}{
		in:  `hello, world`,
		out: `method=count input="hello, world" output=12 err=null`,
	}

	svc := stringsvc.NewStringService()
	ctx := context.Background()

	var buf bytes.Buffer
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(&buf)
		if debug {
			logger = log.NewLogfmtLogger(io.MultiWriter(&buf, os.Stderr))
		}
	}

	logMw := stringsvc.LoggingMiddleware(logger)(svc)
	logMw.Count(ctx, c.in)
	testLogFmt(t, buf, c.out)
}

func testLogFmt(t *testing.T, logOutput bytes.Buffer, prefix string) {
	const logfmtRegex = `^method=(\S+) ` +
		`input=("*.*"*) ` +
		`output=("*.*"*) ` +
		`err=("*.*"*) ` +
		`took=(\S+)$`

	// read log output from buffer
	b, err := ioutil.ReadAll(&logOutput)
	if err != nil {
		t.Error(err)
	}

	// match log output with regex
	match, err := regexp.Match(logfmtRegex, bytes.TrimSpace(b))
	if err != nil {
		t.Error(err)
	}
	if !match {
		t.Errorf("log output does not match regex %s, ouput %s", logfmtRegex, b)
	}

	// check if log output is correct, except for `took=...`
	// since we don't know the exact time our sever serve the request.
	p := []byte(prefix)
	if !bytes.HasPrefix(b, p) {
		t.Errorf("log output method does not match, want prefix %q, got %q", p, b)
	}
}
