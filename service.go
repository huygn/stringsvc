package stringsvc

import (
	"errors"
	"strings"

	"golang.org/x/net/context"
)

// Service provide operations on string.
type Service interface {
	Uppercase(context.Context, string) (string, error)
	Count(context.Context, string) int
}

type stringService struct{}

func NewStringService() Service {
	return stringService{}
}

func (stringService) Uppercase(ctx context.Context, s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(ctx context.Context, s string) int {
	return len(s)
}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty string")
