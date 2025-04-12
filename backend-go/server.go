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
// Hash 3218
// Hash 6628
// Hash 9662
// Hash 6262
// Hash 5586
// Hash 7834
// Hash 5205
// Hash 1358
// Hash 9310
// Hash 8507
// Hash 5813
// Hash 7394
// Hash 9577
// Hash 3036
// Hash 9266
// Hash 5141
// Hash 5945
// Hash 2127
// Hash 4012
// Hash 8353
// Hash 6044
// Hash 9986
// Hash 4372
// Hash 1145
// Hash 9443
// Hash 2377
// Hash 8054
// Hash 5993
// Hash 6975
// Hash 6128
// Hash 5482
// Hash 6445
// Hash 9410
// Hash 3115
// Hash 7110
// Hash 7204
// Hash 5062
// Hash 9711
// Hash 9121
// Hash 8857
// Hash 1887
// Hash 9607
// Hash 9887
// Hash 3391
// Hash 2012
// Hash 1958
// Hash 6870
// Hash 9474
// Hash 6543
// Hash 1741
// Hash 8751
// Hash 3961
// Hash 7082
// Hash 3528
// Hash 1870
// Hash 5831
// Hash 4960
// Hash 2708