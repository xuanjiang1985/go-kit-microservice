package endpoint

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"kit/service"
	"net/http"
)

// AuthRequest
type AuthRequest struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

// AuthResponse
type AuthResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
	Error   string `json:"error"`
}

func MakeAuthEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AuthRequest)

		token, err := svc.Login(ctx, req.Name, req.Pwd)

		var resp AuthResponse
		if err != nil {
			resp = AuthResponse{
				Success: err == nil,
				Token:   token,
				Error:   err.Error(),
			}
		} else {
			resp = AuthResponse{
				Success: err == nil,
				Token:   token,
			}
		}

		return resp, nil
	}
}

func DecodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var loginRequest AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		return nil, err
	}
	return loginRequest, nil
}

func EncodeLoginResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
