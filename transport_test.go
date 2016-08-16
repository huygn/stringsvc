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

func TestHandlers(t *testing.T) {
	cases := []struct {
		method, route    string
		reqBody          string
		expectedRespBody string
		expectedRespCode int
	}{
		{
			method:           "POST",
			route:            "/uppercase",
			reqBody:          `{"s":"hello, world"}`,
			expectedRespBody: `{"v":"HELLO, WORLD"}`,
			expectedRespCode: http.StatusOK,
		},
		{
			method:           "POST",
			route:            "/uppercase",
			reqBody:          `{"s":""}`,
			expectedRespBody: `{"v":"","err":"Empty string"}`,
			expectedRespCode: http.StatusOK,
		},
		{
			method:           "POST",
			route:            "/count",
			reqBody:          `{"s":"hello, world"}`,
			expectedRespBody: `{"v":12}`,
			expectedRespCode: http.StatusOK,
		},
	}

	ctx := context.Background()
	logger := log.NewLogfmtLogger(os.Stderr)
	s := stringsvc.NewStringService()
	h := stringsvc.MakeHTTPHandler(ctx, s, logger)

	for _, c := range cases {
		req, _ := http.NewRequest(c.method, c.route, strings.NewReader(c.reqBody))
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, req)

		errMsg := "%s %s, body: %s - want %v, got %v"

		if rec.Code != c.expectedRespCode {
			t.Errorf(errMsg, c.method, c.route, c.reqBody, c.expectedRespCode, rec.Code)
		}

		respBody := strings.TrimSpace(rec.Body.String())
		if respBody != c.expectedRespBody {
			t.Errorf(errMsg, c.method, c.route, c.reqBody, c.expectedRespBody, respBody)
		}
	}
}
