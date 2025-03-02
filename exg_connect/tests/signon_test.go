package tests

import (
	"log"
	"sync"
	"testing"
	"time"

	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/handlers"
	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/threads"
)

// func TestSendSignonReq(t *testing.T) {
// 	var seq uint32 = 1
// 	sendConn, gatewayInfo, err := handlers.GatewayRouter()
// 	if err != nil {
// 		log.Printf("Error while Gateway Router Request:%+v", err)
// 	}
// 	err = request_handlers.SendSignonReq(sendConn, gatewayInfo, seq)
// 	if err != nil {
// 		log.Printf("Error while SignOn-In :%+v", err)
// 	}
// }

func BenchmarkSendSignonReq(b *testing.B) {
	var seq uint32 = 1
	sendConn, gatewayInfo, err := handlers.GatewayRouter()
	if err != nil {
		log.Printf("Error while Gateway Router Request:%+v", err)
	}
	// defer sendConn.Close() // Close after benchmark
	time.Sleep(1 * time.Second)
	recvConn, err := handlers.BoxRegistration(gatewayInfo)
	if err != nil {
		log.Printf("Error while Box Registration :%+v", err)
	}
	var wg sync.WaitGroup
	wg.Add(1)

	log.Printf("Benchmark started with b.N = %d", b.N)
	b.ResetTimer() // Reset timing after setup
	log.Printf("Benchmark started with b.N = %d", b.N)

	for i := 0; i < b.N; i++ {
		threads.SendToExchange(sendConn, gatewayInfo, seq)
		seq += 1
	}

	go func() {
		// defer wg.Done()
		threads.GetExchangeResp(recvConn, gatewayInfo)
	}()
	wg.Wait()

}
