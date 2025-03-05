package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/agamrai0123/FNO_EXCHANGE/random_order_generator/internal"
	"github.com/agamrai0123/FNO_EXCHANGE/random_order_generator/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// convertToProto converts an internal Order (from your random order generator)
// to a protobuf Order message.
// func convertToProto(order models.Order) *pb.Order {
// 	return &pb.Order{
// 		SessionId:            order.SessionId, // Adjust field names if needed
// 		ExchangeCode:         order.ExchangeCode,
// 		EbaMatchAccount:      order.EbaMatchAccount,
// 		UserId:               order.UserId,
// 		Channel:              order.Channel,
// 		CseId:                order.CseId,
// 		PipeId:               order.PipeId,
// 		CtclId:               order.CtclId,
// 		ProductType:          string(order.ProductType),
// 		Underlying:           order.Underlying,
// 		ExpiryDate:           order.ExpiryDate,
// 		ExcerciseType:        string(order.ExcerciseType),
// 		OptionType:           string(order.OptionType),
// 		StrikePrice:          order.StrikePrice,
// 		IndexOrStock:         string(order.IndexOrStock),
// 		CaLevel:              order.CALevel,
// 		ActionId:             order.ActionId,
// 		BalanceAmount:        order.BalanceAmount,
// 		CanModifyFlag:        string(order.CanModifyFlag),
// 		NkdBlockedFlag:       string(order.NKDBlockedFlag),
// 		ModifyTradeDate:      order.ModifyTradeDate,
// 		ModifyTradeTime:      order.ModifyTradeTime,
// 		SlmFlag:              string(order.SLMFlag),
// 		DisclosedQuantity:    order.DisclosedQuantity,
// 		TotalOrderQuantity:   order.TotalOrderQuantity,
// 		LimitRate:            order.LimitRate,
// 		StopLossTrigger:      order.StopLossTrigger,
// 		OrderValidDate:       order.OrderValidDate,
// 		OrderType:            string(order.OrderType),
// 		AckTime:              order.AckTime,
// 		SpecialFlag:          string(order.SpecialFlag),
// 		OrderFlow:            string(order.OrderFlow),
// 		SpreadOrderIndicator: string(order.SpreadOrderIndicator),
// 		Remarks:              order.Remarks,
// 		UserFlag:             string(order.UserFlag),
// 		ExchangeRemarks:      order.ExchangeRemarks,
// 		IndexCode:            order.IndexCode,
// 		SltpTrailFlag:        string(order.SLTPTrailFlag),
// 		VendorId:             order.VendorId,
// 		VendorNumber:         order.VendorNumber,
// 		OneClickFlag:         string(order.OneClickFlag),
// 		OneClickPortfolioId:  order.OneClickPortfolioId,
// 		AlgoId:               order.AlgoId,
// 		AlgoOrderRemarks:     order.AlgoOrderRemarks,
// 		SourceFlag:           string(order.SourceFlag),
// 		PopupFlag:            string(order.PopupFlag),
// 		ExpiryDate2:          order.ExpiryDate2,
// 		IpAddress:            order.IpAddress,
// 		CallSource:           order.CallSource,
// 		FreshOrderRef:        order.FreshOrderRef,
// 		Alias:                order.Alias,
// 		SystemMessage:        order.SystemMessage,
// 		RequestType:          string(order.RequestType),
// 		UserPassword:         order.UserPassword,
// 		DeliveryEosFlag:      string(order.DeliveryEOSFlag),
// 		OrderReference:       order.OrderReference,
// 		CoverOrderRef:        order.CoverOrderRef,
// 	}
// }

func main() {
	// Connect to the Ingest gRPC server.
	// conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("Failed to connect: %v", err)
	// }
	// defer conn.Close()

	// client := pb.NewIngestClient(conn)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// now := time.Now()
	// log.Println("Start Time: ", now)

	// var wg sync.WaitGroup
	numOrders := 1000
	// wg.Add(numOrders)
	// for range numOrders {
	// 	go internal.GenerateRandomOrder()
	// }
	// wg.Wait()
	for range numOrders {
		order := internal.GenerateRandomOrder()
		if err := sendOrder(order); err != nil {
			fmt.Println("Error:", err)
		}
	}
	// diff := time.Since(now)
	// log.Println("End Time: ", time.Now())
	// log.Printf("Total Time taken: %v\n", diff)
}

// func generateOrder(client pb.IngestClient, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	o := internal.GenerateRandomOrder()
// 	protoOrder := convertToProto(o)
// 	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
// 	defer cancel()

// 	_, err := client.SendOrder(ctx, protoOrder)
// 	if err != nil {
// 		log.Printf("Error sending order %v: %v", o.SessionId, err)
// 		return
// 	}
// 	// log.Printf("Response for order %v: success=%v, message=%s\n", o.SessionId, resp.Success, resp.Message)
// }

func sendOrder(order models.Order) error {
	url := "http://localhost:8080/orders" // Backend URL

	jsonData, err := json.Marshal(order)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send order: %s", resp.Status)
	}

	fmt.Println("Order sent successfully:", order)
	return nil
}
