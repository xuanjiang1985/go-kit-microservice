package service

import (
	"context"
	"errors"
	"strings"
)

type StringService interface {
	Uppercase(context.Context, string) (string, error)
	Count(context.Context, string) int
	Login(context.Context, string, string) (string, error)
}

type BasicStringService struct{}

var ErrEmpty = errors.New("empty string")

func (BasicStringService) Uppercase(_ context.Context, s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (BasicStringService) Count(_ context.Context, s string) int {
	return len(s)
}

func (BasicStringService) Login(_ context.Context, name string, pwd string) (string, error) {
	if name == "zhougang" && pwd == "123abc" {
		token, err := Sign(name, "251")
		return token, err
	}

	return "", errors.New("Your name or password dismatch")
}
