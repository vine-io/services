package server

import (
	"context"

	"github.com/vine-io/vine"
	log "github.com/vine-io/vine/lib/logger"

	"github.com/vine-io/services/pkg/runtime"
	"github.com/vine-io/services/pkg/apiserver/service"
	"github.com/vine-io/services/pkg/runtime/inject"
	pb "github.com/vine-io/services/proto/service/apiserver/v1"
)

type server struct{
	vine.Service

	H service.Apiserver `inject:""`
}

// Call is a single request handler called via client.Call or the generated client code
func (s *server) Call(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	// TODO: Validate
	s.H.Call()
	// FIXME: fix call method
	log.Info("Received Apiserver.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (s *server) Stream(ctx context.Context, req *pb.StreamingRequest, stream pb.ApiserverService_StreamStream) error {
	log.Infof("Received Apiserver.Stream request with count: %d", req.Count)

	// TODO: Validate
	s.H.Stream()
	// FIXME: fix stream method

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&pb.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (s *server) PingPong(ctx context.Context, stream pb.ApiserverService_PingPongStream) error {
	// TODO: Validate
	s.H.PingPong()
	// FIXME: fix stream pingpong

	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
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
