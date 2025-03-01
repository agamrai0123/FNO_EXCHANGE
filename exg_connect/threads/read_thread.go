package threads

import (
	"encoding/binary"
	"io"
	"log"
	"net"

	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/models"
	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/response_handlers"
	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/utils"
)

func GetExchangeResp(conn net.Conn, GatewayInfo *models.GatewayRouterResponse) {
	for {
		buf := make([]byte, 1024)
		_, err := io.ReadFull(conn, buf)
		if err != nil {
			if err.Error() == "EOF" {
				// log.Println("nothing to read")
				continue
			}
			// log.Println("failed to read:", err)
			// break
		} else {
			log.Print("data recieved")
			exgdata := &models.ExchangeData{
				Length:         binary.LittleEndian.Uint16(buf[0:2]),
				SequenceNumber: binary.LittleEndian.Uint32(buf[2:6]),
				Checksum:       buf[6:22],
				MessageData:    buf[22:],
			}
			decrypted, err := utils.DecryptAES(GatewayInfo.CryptographicKey, GatewayInfo.CryptographicIV, exgdata.MessageData)
			if err != nil {
				log.Println("failed to decrypt:", err)
			}
			MessageHeader, err := response_handlers.GetHeader(conn, decrypted[0:40])
			if err != nil {
				log.Println("failed to read:", err)
				continue
			}
			ReadRemaining := int(MessageHeader.MessageLength) - 40
			if ReadRemaining > 0 {
				if MessageHeader.TransactionCode == 23001 {
					boxsignonreq, err := response_handlers.ReadBoxSignOnResp(conn, MessageHeader, decrypted)
					if err != nil {
						log.Println("failed to read for transcode", MessageHeader.TransactionCode, err)
						continue
					}
					log.Print(boxsignonreq)
				} else if MessageHeader.TransactionCode == 2301 {
					signonreq, err := response_handlers.ReadSignOnResp(conn, MessageHeader, decrypted)
					if err != nil {
						log.Println("failed to read for transcode", MessageHeader.TransactionCode, err)
						continue
					}
					log.Print(signonreq)
				} else if MessageHeader.TransactionCode == 1601 {

				} else if MessageHeader.TransactionCode == 7307 {

				}
			}
		}
	}
}
