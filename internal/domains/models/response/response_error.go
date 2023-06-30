package response

import (
	"github.com/HEEPOKE/gin-hexagonal-graphql/pkg/constants"
)

type Response struct {
	Payload interface{}                `json:"payload"`
	Status  *constants.ResponseMessage `json:"status"`
}

func NewResponse(payload interface{}, status *constants.ResponseMessage) *Response {
	return &Response{
		Payload: payload,
		Status:  status,
	}
}
