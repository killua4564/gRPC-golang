package main

import (  
    "log"

    "golang.org/x/net/context"
    "google.golang.org/grpc"

    "gRPC/pb"
)

func main() {  
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("connecting failed: %v", err)
    }
    defer conn.Close()

    c := pb.NewCalculatorClient(conn)

    r, err := c.Plus(context.Background(), &pb.CalcRequest{NumberA: 2, NumberB: 3})
    if err != nil {
        log.Fatalf("plus error: %v", err)
    }
    log.Printf("result: %d", r.Result)
}