package greetings

import (
	"fmt"
	"net/http"

	"github.com/shuryak/api-wrappers/internal/api"
)

type HelloReq struct {
	Name string `json:"name"`
}

func (req HelloReq) Validate(_ *api.Context) error {
	return nil
}

type HelloResp struct {
	Message string `json:"message"`
}

func Hello(_ *api.Context, req *HelloReq) (*HelloResp, int) {
	return &HelloResp{
		Message: fmt.Sprintf("Hello, %s!", req.Name),
	}, http.StatusOK
}
