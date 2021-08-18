// MIT License
//
// Copyright (c) 2021 Lack
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package server

import (
	"github.com/vine-io/services/pkg/apiserver/service"
	"github.com/vine-io/services/pkg/runtime"
	"github.com/vine-io/services/pkg/runtime/inject"
	pb "github.com/vine-io/services/proto/service/apiserver/v1"
	"github.com/vine-io/vine"
)

type server struct {
	vine.Service

	H service.Apiserver `inject:""`
}

func (s *server) Init() error {
	var err error

	opts := []vine.Option{
		vine.Name(runtime.ApiserverName),
		vine.Id(runtime.ApiserverId),
		vine.Version(runtime.GetVersion()),
		vine.Metadata(map[string]string{
			"namespace": runtime.Namespace,
		}),
	}

	s.Service.Init(opts...)

	if err = inject.Provide(s.Service, s.Client(), s); err != nil {
		return err
	}

	// TODO: inject more objects

	if err = inject.Populate(); err != nil {
		return err
	}

	if err = s.H.Init(); err != nil {
		return err
	}

	if err = pb.RegisterApiserverServiceHandler(s.Service.Server(), s); err != nil {
		return err
	}

	return err
}

func New() *server {
	srv := vine.NewService()
	return &server{
		Service: srv,
	}
}
