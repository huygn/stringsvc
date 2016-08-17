package stringsvc_test

import (
	"testing"

	"golang.org/x/net/context"

	"github.com/gnhuy91/stringsvc"
)

func TestUppercase(t *testing.T) {
	cases := []struct {
		in, out string
		err     error
	}{
		{
			in:  "hello, world",
			out: "HELLO, WORLD",
			err: nil,
		},
		{
			in:  "",
			out: "",
			err: stringsvc.ErrEmpty,
		},
	}

	ctx := context.Background()
	svc := stringsvc.NewStringService()

	for _, c := range cases {
		s, err := svc.Uppercase(ctx, c.in)
		if err != c.err {
			t.Errorf("input: %q, want %v, got %v", c.in, c.err, err)
		}
		if s != c.out {
			t.Errorf("input: %q, want %q, got %q", c.in, c.out, s)
		}
	}
}

func TestCount(t *testing.T) {
	c := struct {
		in  string
		out int
	}{
		in:  "hello, world",
		out: 12,
	}

	ctx := context.Background()
	svc := stringsvc.NewStringService()

	n := svc.Count(ctx, c.in)
	if n != c.out {
		t.Errorf("input: %q, want %v, got %v", c.in, c.out, n)
	}
}
