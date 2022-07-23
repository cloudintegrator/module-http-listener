package httplistener

import (
	"context"
	"fmt"
	"github.com/cloudintegrator/cloud-integrator/pkg/listener/spec"
)

type HttpListener struct {
	Metadata spec.ListenerSpec
}

func (listener *HttpListener) Start(ctx *context.Context) {
	fmt.Println("Listener has started listening...")
}
