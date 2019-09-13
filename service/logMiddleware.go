package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"time"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   StringService
}

func (mw LoggingMiddleware) Uppercase(ctx context.Context, s string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.Uppercase(ctx, s)
	return
}

func (mw LoggingMiddleware) Count(ctx context.Context, s string) (n int) {
	//defer func(begin time.Time) {
	//	_ = mw.Logger.Log(
	//		"method", "count",
	//		"input", s,
	//		"n", n,
	//		"took", time.Since(begin),
	//	)
	//}(time.Now())
	//
	//n = mw.Next.Count(s)

	num := mw.Next.Count(ctx, s)
	_ = mw.Logger.Log(
		"method", "count",
		"input", s,
		"n", num,
		"time", time.Now().UnixNano(),
	)

	return num
}

func (mw LoggingMiddleware) Login(ctx context.Context, name string, pwd string) (output string, err error) {
	//defer func(begin time.Time) {
	//	_ = mw.Logger.Log(
	//		"method", "count",
	//		"input", s,
	//		"n", n,
	//		"took", time.Since(begin),
	//	)
	//}(time.Now())
	//
	//n = mw.Next.Count(s)

	output, err = mw.Next.Login(ctx, name, pwd)
	_ = mw.Logger.Log(
		"method", "login",
		"input", name,
		"pwd", pwd,
		"time", time.Now().UnixNano(),
	)

	return output, err
}
