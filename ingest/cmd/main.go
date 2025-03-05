package main

import (
	"log"
	"net/http"
	"time"

	"github.com/agamrai0123/FNO_EXCHANGE/ingest/models"
	"github.com/agamrai0123/FNO_EXCHANGE/ingest/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var ValidOrders uint16

// ingestServer implements the generated IngestServer interface.
// type ingestServer struct {
// 	pb.UnimplementedIngestServer
// }

// func (s *ingestServer) SendOrder(ctx context.Context, pborder *pb.Order) (*pb.OrderResponse, error) {

// 	Order, err := utils.ConvertProtoToModel(pborder)
// 	if err != nil {
// 		log.Printf("Conversion error: %v", err)
// 		return &pb.OrderResponse{Success: false, Message: "Conversion failed"}, nil
// 	}
// 	err = utils.ValidateOrderInputs(&Order)
// 	if err != nil {
// 		log.Printf("Validation error: %v", err)
// 		return &pb.OrderResponse{Success: false, Message: "Validation failed"}, nil
// 	}
// 	ValidOrders++
// 	// log.Printf("Received order: %+v", Order)
// 	log.Printf("%d valid orders", ValidOrders)
// 	// Process the order as needed.
// 	return &pb.OrderResponse{Success: true, Message: "Order processed successfully"}, nil
// }

// var i = 0

func main() {
	// lis, err := net.Listen("tcp", ":50051")
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// grpcServer := grpc.NewServer()
	// pb.RegisterIngestServer(grpcServer, &ingestServer{})

	// log.Println("Ingest service is running on port 50051")
	// if err := grpcServer.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type"},
		MaxAge:       24 * time.Hour,
	}))

	// Define API endpoint to receive orders
	r.POST("/orders", createOrder)

	// Start server on port 8080
	r.Run(":8080")

}

func createOrder(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order format"})
		return
	}

	// Validate order
	if err := utils.ValidateOrderInputs(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order format"})
		return
	}
	start := time.Now()
	// Send order to Kafka for processing
	if err := utils.Producer(&order); err != nil {
		log.Println("Kafka error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process order"})
		return
	}
	log.Println("Validation time:", time.Since(start))

	c.JSON(http.StatusOK, gin.H{"message": "Order received"})
}

// func sendToKafka(order *models.Order) error {
// 	var topic string
// 	if order.ProductType == 'O' {
// 		topic = "Options"
// 	} else {
// 		return errors.New("topic not specified")
// 	}
// 	writer := kafka.Writer{
// 		Addr:     kafka.TCP("localhost:9092"),
// 		Topic:    topic,
// 		Balancer: &kafka.LeastBytes{},
// 	}

// 	orderBytes, err := json.Marshal(order)
// 	if err != nil {
// 		return err
// 	}

// 	err = writer.WriteMessages(context.Background(),
// 		kafka.Message{
// 			Value: orderBytes,
// 		})
// 	if err != nil {
// 		return err
// 	}

// 	log.Println("Order pushed to Kafka")
// 	return nil
// }
