package calc

import (
	"errors"
	"net/http"

	"github.com/shuryak/api-wrappers/internal/api"
)

type CalculateReq struct {
	FirstNumber   int    `json:"first_number"`
	SecondNumber  int    `json:"second_number"`
	OperationName string `json:"operation_name"`
}

func (req CalculateReq) Validate(_ *api.Context) error {
	if req.FirstNumber < 0 {
		return errors.New("first_number < 0")
	}
	if req.SecondNumber < 0 {
		return errors.New("second_number < 0")
	}
	if req.OperationName != "plus" && req.OperationName != "multiply" {
		return errors.New("invalid operator_name")
	}

	return nil
}

type CalculateResp struct {
	Result int `json:"result"`
}

func (c *Calc) Handle(_ *api.Context, req *CalculateReq) (*CalculateResp, int) {
	resp := &CalculateResp{}

	switch req.OperationName {
	case "plus":
		resp.Result = req.FirstNumber + req.SecondNumber
	case "multiply":
		resp.Result = req.FirstNumber * req.SecondNumber
	}

	return resp, http.StatusOK
}
