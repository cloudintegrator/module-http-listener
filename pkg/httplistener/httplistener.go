package httplistener

import (
	"context"
	"fmt"
	"github.com/cloudintegrator/module-http-listener/pkg/httplistener/spec"
)

type HttpListener struct {
	Config spec.HttpListenerConfig `json:"config"`
}

func (listener *HttpListener) Start(ctx *context.Context) {
	fmt.Println("Listener has started listening...")
}
