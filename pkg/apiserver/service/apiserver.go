package service

import (
	"github.com/vine-io/services/pkg/runtime/inject"
	"github.com/vine-io/vine"
)

func init() {
	inject.ProvidePanic(new(apiserver))
}

type Apiserver interface {
	Init() error
	Call()
	Stream()
	PingPong()
}

var _ Apiserver = (*apiserver)(nil)

type apiserver struct {
	vine.Service `inject:""`
}

func (s *apiserver) Init() error {
	return nil
}

func (s *apiserver) Call() {
	// FIXME: modify method
	panic("implement me")
}

func (s *apiserver) Stream() {
	panic("implement me")
}

func (s *apiserver) PingPong() {
	panic("implement me")
}
