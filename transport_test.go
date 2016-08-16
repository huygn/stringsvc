package stringsvc_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"

	"github.com/gnhuy91/stringsvc"
)

func TestHandler(t *testing.T) {
	cases := []struct {
		method, url      string
		reqBody          string
		expectedRespBody string
		expectedRespCode int
	}{
		{
			method:           "POST",
			url:              "/uppercase",
			reqBody:          `{"s":"hello, world"}`,
			expectedRespBody: `{"v":"HELLO, WORLD"}`,
			expectedRespCode: http.StatusOK,
		},
		{
			method:           "POST",
			url:              "/count",
			reqBody:          `{"s":"hello, world"}`,
			expectedRespBody: `{"v":12}`,
			expectedRespCode: http.StatusOK,
		},
	}

	for _, c := range cases {
		ctx := context.Background()
		logger := log.NewLogfmtLogger(os.Stderr)
		s := stringsvc.NewStringService()
		h := stringsvc.MakeHTTPHandler(ctx, s, logger)

		req, _ := http.NewRequest(c.method, c.url, strings.NewReader(c.reqBody))
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		errMsg := "%s %s, body: %s - want %v, got %v"

		if rec.Code != c.expectedRespCode {
			t.Errorf(errMsg, c.method, c.url, c.reqBody, c.expectedRespCode, rec.Code)
		}

		respBody := strings.TrimSpace(rec.Body.String())
		if respBody != c.expectedRespBody {
			t.Errorf(errMsg, c.method, c.url, c.reqBody, c.expectedRespBody, respBody)
		}
	}
}
