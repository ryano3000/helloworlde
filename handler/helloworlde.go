package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	helloworlde "helloworlde/proto"
)

type Helloworlde struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Helloworlde) Call(ctx context.Context, req *helloworlde.Request, rsp *helloworlde.Response) error {
	log.Info("Received Helloworlde.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Helloworlde) Stream(ctx context.Context, req *helloworlde.StreamingRequest, stream helloworlde.Helloworlde_StreamStream) error {
	log.Infof("Received Helloworlde.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&helloworlde.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Helloworlde) PingPong(ctx context.Context, stream helloworlde.Helloworlde_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&helloworlde.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
