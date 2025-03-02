package main

import (
	"log"
	"sync"
	"time"

	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/handlers"
	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/request_handlers"
	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/threads"
)

func main() {
	var seq uint32 = 1
	sendConn, gatewayInfo, err := handlers.GatewayRouter()
	if err != nil {
		log.Printf("Error while Gateway Router Request:%+v", err)
	}

	log.Println("Connected to server:", sendConn.RemoteAddr())

	time.Sleep(1 * time.Second)
	recvConn, err := handlers.BoxRegistration(gatewayInfo)
	if err != nil {
		log.Printf("Error while Box Registration :%+v", err)
	}
	// defer conn.Close()

	log.Println("Connected to server:", recvConn.RemoteAddr())
	// time.Sleep(5 * time.Second)
	err = request_handlers.SendBoxSignonReq(sendConn, gatewayInfo, seq)
	if err != nil {
		log.Printf("Error while Box Registration :%+v", err)
	}

	// time.Sleep(200 * time.Second)
	var wg sync.WaitGroup
	wg.Add(1)

	threads.SendToExchange(sendConn, gatewayInfo, seq)
	go func() {
		defer wg.Done()
		threads.GetExchangeResp(recvConn, gatewayInfo)
	}()
	wg.Wait()

}
