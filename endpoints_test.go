package stringsvc_test

import (
	"fmt"
	"reflect"
	"testing"

	"golang.org/x/net/context"

	"github.com/gnhuy91/stringsvc"
)

func TestUppercaseEnpoint(t *testing.T) {
	cases := []struct {
		in, out, err string
	}{
		{
			in:  "hello, world",
			out: "HELLO, WORLD",
			err: "",
		},
		{
			in:  "",
			out: "",
			err: "Empty string",
		},
	}

	ctx := context.Background()
	e := stringsvc.MakeUppercaseEndpoint(stringsvc.NewStringService())

	for _, c := range cases {
		t.Run(
			fmt.Sprintf("in=%q", c.in),
			func(t *testing.T) {
				req := stringsvc.NewTestUppercaseRequest(c.in)
				expectedResp := stringsvc.NewTestUppercaseResponse(c.out, c.err)

				resp, _ := e(ctx, req)
				eq := reflect.DeepEqual(resp, expectedResp)
				if !eq {
					t.Errorf("%q => %+v, want %+v", c.in, resp, expectedResp)
				}
			})
	}
}

func TestCountEnpoint(t *testing.T) {
	c := struct {
		in  string
		out int
	}{

		in:  "hello, world",
		out: 12,
	}

	ctx := context.Background()
	e := stringsvc.MakeCountEndpoint(stringsvc.NewStringService())

	req := stringsvc.NewTestCountRequest(c.in)
	expectedResp := stringsvc.NewTestCountResponse(c.out)

	resp, _ := e(ctx, req)
	eq := reflect.DeepEqual(resp, expectedResp)
	if !eq {
		t.Errorf("%q => %+v, want %+v", c.in, resp, expectedResp)
	}
}
