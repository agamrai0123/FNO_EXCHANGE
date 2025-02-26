package main

import (
	"log"

	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/handlers"
	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/request_handlers"
	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/threads"
)

func main() {
	var seq uint32 = 1
	gatewayInfo, err := handlers.GatewayRouter()
	if err != nil {
		log.Printf("Error while Gateway Router Request:%+v", err)
	}

	conn, err := handlers.BoxRegistration(gatewayInfo)
	if err != nil {
		log.Printf("Error while Box Registration :%+v", err)
	}
	// defer conn.Close()

	log.Println("Connected to server:", conn.RemoteAddr())

	err = request_handlers.SendBoxSignonReq(conn, gatewayInfo, seq)
	if err != nil {
		log.Printf("Error while Box Registration :%+v", err)
	}

	err = request_handlers.SendSignonReq(conn, gatewayInfo, seq)
	if err != nil {
		log.Printf("Error while SignOn-In :%+v", err)
	}

	go threads.GetExchangeResp(conn, gatewayInfo)
	// go threads.SendToExchange(conn)
}
