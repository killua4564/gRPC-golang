package main

import (  
    "log"
    "net"

    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"

    "gRPC/pb"
)

type server struct{}

func (s *server) Plus(ctx context.Context, in *pb.CalcRequest) (*pb.CalcReply, error) {
    result := in.NumberA + in.NumberB
    return &pb.CalcReply{Result: result}, nil
}

func main() {  
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Listen port fail: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterCalculatorServer(s, &server{})

    reflection.Register(s)

    if err := s.Serve(lis); err != nil {
        log.Fatalf("Serve Error: %v", err)
    }
}