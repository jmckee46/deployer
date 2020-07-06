package responses

import (
	"net/http"

	"github.com/jmckee46/deployer/serializers"
)

type Response struct {
	Head           interface{}            `json:"head,omitempty"`
	Body           interface{}            `json:"body,omitempty"`
	Writer         http.ResponseWriter    `json:"-"`
	HTTPStatusCode int                    `json:"http-status-code"`
	Serializer     serializers.Serializer `json:"-"`
	Sent           bool                   `json:"sent"`
}

func NewResponse() *Response {
	return &Response{}
}
