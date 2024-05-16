package greetings

import (
	"fmt"
	"net/http"

	"github.com/shuryak/api-wrappers/internal/api"
)

type HelloGetReq struct {
	Name string `query:"name"`
}

func (req HelloGetReq) Validate(_ *api.Context) error {
	return nil
}

type HelloGetResp struct {
	Message string `json:"message"`
}

func HelloGet(_ *api.Context, req *HelloGetReq) (*HelloGetResp, int) {
	return &HelloGetResp{
		Message: fmt.Sprintf("Hello, %s!", req.Name),
	}, http.StatusOK
}
