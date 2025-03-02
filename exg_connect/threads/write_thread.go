package threads

import (
	"log"
	"net"

	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/models"
	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/request_handlers"
)

func SendToExchange(conn net.Conn, gatewayInfo *models.GatewayRouterResponse, seq uint32) {
	err := request_handlers.SendSignonReq(conn, gatewayInfo, seq)
	if err != nil {
		log.Printf("Error while SignOn-In :%+v", err)
	}
}
