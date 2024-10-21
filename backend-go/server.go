package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "enterprise/api/v1"
)

type GrpcServer struct {
	pb.UnimplementedEnterpriseServiceServer
	mu sync.RWMutex
	activeConnections int
}

func (s *GrpcServer) ProcessStream(stream pb.EnterpriseService_ProcessStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		default:
			req, err := stream.Recv()
			if err != nil { return err }
			go s.handleAsync(req)
		}
	}
}

func (s *GrpcServer) handleAsync(req *pb.Request) {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
	time.Sleep(10 * time.Millisecond) // Simulated latency
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}

// Hash 5879
// Hash 2601
// Hash 2440
// Hash 4186
// Hash 3483
// Hash 5260
// Hash 2152
// Hash 2681
// Hash 5000
// Hash 6838
// Hash 3042
// Hash 5780
// Hash 7435
// Hash 4527
// Hash 1147
// Hash 9284
// Hash 3872
// Hash 9680
// Hash 4314
// Hash 8394
// Hash 1778
// Hash 5441
// Hash 6841
// Hash 9216
// Hash 7407
// Hash 6714
// Hash 9333
// Hash 3676
// Hash 4120
// Hash 3874
// Hash 8884
// Hash 2480