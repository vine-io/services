package handler

import (
	"context"

	"github.com/vine-io/services/auth"
	"github.com/vine-io/vine"
	log "github.com/vine-io/vine/lib/logger"

	pb "github.com/vine-io/services/api/service/auth/v1"
)

type server struct {
	vine.Service
}

// Call is a single request handler called via client.Call or the generated client code
func (s *server) Call(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	// TODO: Validate
	// FIXME: fix call method
	log.Info("Received Auth.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (s *server) Stream(ctx context.Context, req *pb.StreamingRequest, stream pb.AuthService_StreamStream) error {
	log.Infof("Received Auth.Stream request with count: %d", req.Count)

	// TODO: Validate
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
func (s *server) PingPong(ctx context.Context, stream pb.AuthService_PingPongStream) error {
	// TODO: Validate
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
		vine.Name(auth.Name),
		vine.Id(auth.Id),
		vine.Version(auth.GetVersion()),
		vine.Metadata(map[string]string{
			"namespace": auth.Namespace,
		}),
	}

	s.Service.Init(opts...)

	if err = pb.RegisterAuthServiceHandler(s.Service.Server(), s); err != nil {
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
