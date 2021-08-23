package service

import (
	"github.com/vine-io/services/pkg/runtime/inject"
	"github.com/vine-io/vine"
)

func init() {
	inject.ProvidePanic(new(auth))
}

type Auth interface {
	Init() error
	Call()
	Stream()
	PingPong()
}

var _ Auth = (*auth)(nil)

type auth struct {
	vine.Service `inject:""`
}

func (s *auth) Init() error {
	return nil
}

func (s *auth) Call() {
	// FIXME: modify method
	panic("implement me")
}

func (s *auth) Stream() {
	panic("implement me")
}

func (s *auth) PingPong() {
	panic("implement me")
}
