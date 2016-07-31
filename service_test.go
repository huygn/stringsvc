package stringsvc

import (
	"testing"

	"golang.org/x/net/context"
)

func TestUppercase(t *testing.T) {
	const (
		inp  = "hello, world"
		outp = "HELLO, WORLD"
	)

	ctx := context.Background()
	svc := NewStringService()

	s, err := svc.Uppercase(ctx, inp)
	if err != nil {
		t.Errorf(err.Error())
	}
	if s != outp {
		t.Errorf("input: %s, want %s, got %s", inp, outp, s)
	}
}

func TestUppercase_FailIfInputNil(t *testing.T) {
	const (
		inp = ""
	)

	ctx := context.Background()
	svc := NewStringService()

	_, err := svc.Uppercase(ctx, inp)
	if err != ErrEmpty {
		t.Errorf("input: %s, want %s, got %s", inp, ErrEmpty.Error(), err.Error())
	}
}

func TestCount(t *testing.T) {
	const (
		inp  = "hello, world"
		outp = 12
	)

	ctx := context.Background()
	svc := NewStringService()

	n := svc.Count(ctx, inp)
	if n != outp {
		t.Errorf("input: %s, want %v, got %v", inp, outp, n)
	}
}
