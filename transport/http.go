package transport

import (
	"github.com/dgrijalva/jwt-go"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/log"
	httpTransport "github.com/go-kit/kit/transport/http"
	"kit/endpoint"
	"kit/service"
	"net/http"
	"os"
)

func Run() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc service.StringService
	svc = service.BasicStringService{}
	svc = service.LoggingMiddleware{logger, svc}

	jwtOptions := []httpTransport.ServerOption{
		httpTransport.ServerBefore(kitjwt.HTTPToContext()),
	}

	uppercaseHandler := httpTransport.NewServer(
		kitjwt.NewParser(service.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(endpoint.MakeUppercaseEndpoint(svc)),
		//endpoint.MakeUppercaseEndpoint(svc),
		endpoint.DecodeUpppercaseRequest,
		endpoint.EncodeResponse,
		jwtOptions...,
	)

	countHandler := httpTransport.NewServer(
		endpoint.MakeCountEndpoint(svc),
		endpoint.DecodeCountRequest,
		endpoint.EncodeResponse,
	)

	var authSvc service.StringService
	authSvc = service.BasicStringService{}
	authSvc = service.LoggingMiddleware{logger, authSvc}

	loginHandle := httpTransport.NewServer(
		endpoint.MakeAuthEndpoint(authSvc),
		endpoint.DecodeLoginRequest,
		endpoint.EncodeLoginResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	http.Handle("/login", loginHandle)
	logger.Log(http.ListenAndServe(":8081", nil))
}
