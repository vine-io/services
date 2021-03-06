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

package handler

import (
	"context"
	"strings"

	pb "github.com/vine-io/services/api/service/apiserver/v1"
	"google.golang.org/grpc/peer"
)

func (h *handler) Healthz(ctx context.Context, _ *pb.Empty, _ *pb.Empty) error {
	return nil
}

func (h *handler) GetIP(ctx context.Context, _ *pb.Empty, rsp *pb.IPRsp) error {
	pr, ok := peer.FromContext(ctx)
	if ok {
		addr := pr.Addr.String()
		if index := strings.LastIndex(addr, ":"); index > 1 {
			addr = addr[0:index]
		}
		rsp.Addr = addr
	}
	return nil
}
