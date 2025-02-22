package main

import (
	"context"
	"log"
	"net"

	"github.com/agamrai0123/FNO_EXCHANGE/ingest/utils"
	pb "github.com/agamrai0123/FNO_EXCHANGE/proto"
	"google.golang.org/grpc"
)

var ValidOrders uint16

// ingestServer implements the generated IngestServer interface.
type ingestServer struct {
	pb.UnimplementedIngestServer
}

func (s *ingestServer) SendOrder(ctx context.Context, pborder *pb.Order) (*pb.OrderResponse, error) {

	Order, err := utils.ConvertProtoToModel(pborder)
	if err != nil {
		log.Printf("Conversion error: %v", err)
		return &pb.OrderResponse{Success: false, Message: "Conversion failed"}, nil
	}
	err = utils.ValidateOrderInputs(&Order)
	if err != nil {
		log.Printf("Validation error: %v", err)
		return &pb.OrderResponse{Success: false, Message: "Validation failed"}, nil
	}
	ValidOrders++
	// log.Printf("Received order: %+v", Order)
	log.Printf("%d valid orders", ValidOrders)
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
