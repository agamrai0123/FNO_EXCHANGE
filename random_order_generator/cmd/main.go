package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/agamrai0123/FNO_EXCHANGE/proto"
	"github.com/agamrai0123/FNO_EXCHANGE/random_order_generator/internal"
	"github.com/agamrai0123/FNO_EXCHANGE/random_order_generator/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// convertToProto converts an internal Order (from your random order generator)
// to a protobuf Order message.
func convertToProto(order models.Order) *pb.Order {
	return &pb.Order{
		SessionId:            order.SessionId, // Adjust field names if needed
		ExchangeCode:         order.ExchangeCode,
		EbaMatchAccount:      order.EbaMatchAccount,
		UserId:               order.UserId,
		Channel:              order.Channel,
		CseId:                order.CseId,
		PipeId:               order.PipeId,
		CtclId:               order.CtclId,
		ProductType:          string(order.ProductType),
		Underlying:           order.Underlying,
		ExpiryDate:           order.ExpiryDate,
		ExcerciseType:        string(order.ExcerciseType),
		OptionType:           string(order.OptionType),
		StrikePrice:          order.StrikePrice,
		IndexOrStock:         string(order.IndexOrStock),
		CaLevel:              order.CALevel,
		ActionId:             order.ActionId,
		BalanceAmount:        order.BalanceAmount,
		CanModifyFlag:        string(order.CanModifyFlag),
		NkdBlockedFlag:       string(order.NKDBlockedFlag),
		ModifyTradeDate:      order.ModifyTradeDate,
		ModifyTradeTime:      order.ModifyTradeTime,
		SlmFlag:              string(order.SLMFlag),
		DisclosedQuantity:    order.DisclosedQuantity,
		TotalOrderQuantity:   order.TotalOrderQuantity,
		LimitRate:            order.LimitRate,
		StopLossTrigger:      order.StopLossTrigger,
		OrderValidDate:       order.OrderValidDate,
		OrderType:            string(order.OrderType),
		AckTime:              order.AckTime,
		SpecialFlag:          string(order.SpecialFlag),
		OrderFlow:            string(order.OrderFlow),
		SpreadOrderIndicator: string(order.SpreadOrderIndicator),
		Remarks:              order.Remarks,
		UserFlag:             string(order.UserFlag),
		ExchangeRemarks:      order.ExchangeRemarks,
		IndexCode:            order.IndexCode,
		SltpTrailFlag:        string(order.SLTPTrailFlag),
		VendorId:             order.VendorId,
		VendorNumber:         order.VendorNumber,
		OneClickFlag:         string(order.OneClickFlag),
		OneClickPortfolioId:  order.OneClickPortfolioId,
		AlgoId:               order.AlgoId,
		AlgoOrderRemarks:     order.AlgoOrderRemarks,
		SourceFlag:           string(order.SourceFlag),
		PopupFlag:            string(order.PopupFlag),
		ExpiryDate2:          order.ExpiryDate2,
		IpAddress:            order.IpAddress,
		CallSource:           order.CallSource,
		FreshOrderRef:        order.FreshOrderRef,
		Alias:                order.Alias,
		SystemMessage:        order.SystemMessage,
		RequestType:          string(order.RequestType),
		UserPassword:         order.UserPassword,
		DeliveryEosFlag:      string(order.DeliveryEOSFlag),
		OrderReference:       order.OrderReference,
		CoverOrderRef:        order.CoverOrderRef,
	}
}

func main() {
	// Connect to the Ingest gRPC server.
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewIngestClient(conn)

	// Generate and send a single random order.
	order := internal.GenerateRandomOrder()
	fmt.Printf("Generated single order: %+v\n", order)
	protoOrder := convertToProto(order)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.SendOrder(ctx, protoOrder)
	if err != nil {
		log.Fatalf("Error sending order: %v", err)
	}
	fmt.Printf("Response for single order: success=%v, message=%s\n", resp.Success, resp.Message)

	// Generate and send multiple random orders.
	orders := internal.GenerateRandomOrders(100)
	fmt.Printf("Generated %d orders\n", len(orders))
	for _, o := range orders {
		protoOrder := convertToProto(o)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		// Note: Call cancel in each loop iteration to avoid context leaks.
		defer cancel()
		resp, err := client.SendOrder(ctx, protoOrder)
		if err != nil {
			log.Printf("Error sending order %v: %v", o.SessionId, err)
			continue
		}
		fmt.Printf("Response for order %v: success=%v, message=%s\n", o.SessionId, resp.Success, resp.Message)
	}
}
