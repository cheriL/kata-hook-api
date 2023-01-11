package kata_hook_api

import (
	"fmt"
	"log"
)

const(
	preStart  = "prestart"
	postStart = "poststart"
	postStop  = "poststop"
)

type controller struct {
	option    string
	handlers  HookHandlers
	cfg       *config
}



func (c *controller) Execute() error {
	log.Printf("Execute() Option: %s", c.option)
	switch c.option {
	case preStart:
		c.handlers.DoPreStart(c.cfg)
	case postStart:
		c.handlers.DoPostStart(c.cfg)
	case postStop:
		c.handlers.DoPostStop(c.cfg)
	default:
		return fmt.Errorf("not support option: %s", c.option)
	}

	return nil
}

type HookHandlers struct {
	PreStart  func(obj interface{})
	PostStart func(obj interface{})
	PostStop  func(obj interface{})
}

func (h HookHandlers) DoPreStart(obj interface{}) {
	if h.PreStart != nil {
		h.PreStart(obj)
	}
}

func (h HookHandlers) DoPostStart(obj interface{}) {
	if h.PostStart != nil {
		h.PostStart(obj)
	}
}

func (h HookHandlers) DoPostStop(obj interface{}) {
	if h.PreStart != nil {
		h.PostStop(obj)
	}
}