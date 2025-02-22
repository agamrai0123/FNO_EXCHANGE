package main

import (
	"context"
	"log"
	"net"

	pb "github.com/agamrai0123/FNO_EXCHANGE/proto"
	"google.golang.org/grpc"
)

// ingestServer implements the generated IngestServer interface.
type ingestServer struct {
	pb.UnimplementedIngestServer
}

func (s *ingestServer) SendOrder(ctx context.Context, order *pb.Order) (*pb.OrderResponse, error) {
	log.Printf("Received order: %+v", order)
	// Process the order as needed.
	return &pb.OrderResponse{Success: true, Message: "Order processed successfully"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterIngestServer(grpcServer, &ingestServer{})

	log.Println("Ingest service is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
