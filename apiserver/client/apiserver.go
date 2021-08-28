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

package client

import (
	"context"

	pb "github.com/vine-io/services/api/service/apiserver/v1"
	"github.com/vine-io/services/apiserver"
	"github.com/vine-io/vine/core/client"
)

type Client interface {
	Healthz(ctx context.Context, opts ...client.CallOption) error
	GetIP(ctx context.Context, opts ...client.CallOption) (string, error)
}

type SimpleClient struct {
	cc pb.APIServerService
}

func NewClient(conn client.Client) *SimpleClient {
	return &SimpleClient{
		cc: pb.NewAPIServerService(apiserver.Name, conn),
	}
}

func (c *SimpleClient) Healthz(ctx context.Context, opts ...client.CallOption) error {
	_, err := c.cc.Healthz(ctx, &pb.Empty{}, opts...)
	return err
}

func (c *SimpleClient) GetIP(ctx context.Context, opts ...client.CallOption) (string, error) {
	rsp, err := c.cc.GetIP(ctx, &pb.Empty{}, opts...)
	if err != nil {
		return "", err
	}
	return rsp.Addr, nil
}
