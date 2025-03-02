package threads

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"io"
	"log"
	"net"

	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/models"
	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/response_handlers"
	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/utils"
)

func GetExchangeResp(conn net.Conn, GatewayInfo *models.GatewayRouterResponse) {

	log.Print("in ReadThread")
	for {
		// Read the length from socket
		lengthBuf := make([]byte, 2)
		_, err := io.ReadFull(conn, lengthBuf)
		if err != nil {
			if err == io.EOF {
				log.Println("Connection closed by peer")
				break
			}
			log.Println("failed to read message length:", err)
			break
		}
		// Convert length to uint16
		messageLength := binary.LittleEndian.Uint16(lengthBuf)
		// log.Printf("Expected message length: %d", messageLength)

		// Step 2: Read the remaining message
		buf := make([]byte, messageLength+16)
		_, err = io.ReadFull(conn, buf)
		if err != nil {
			log.Println("failed to read full message:", err)
			break
		}
		exgdata := &models.ExchangeData{
			Length:         messageLength,
			SequenceNumber: binary.LittleEndian.Uint32(buf[0:4]),
			Checksum:       buf[4:20],
			MessageData:    buf[20:],
		}

		// Decryption of the recieved packet
		decrypted, err := utils.DecryptAES(GatewayInfo.CryptographicKey, GatewayInfo.CryptographicIV, exgdata.MessageData)
		if err != nil {
			log.Println("failed to decrypt:", err)
		}

		// Verify md5 Checksum
		respChecksum := md5.Sum(decrypted)
		if bytes.Equal(exgdata.Checksum, respChecksum[:]) {
			log.Println("md5 checksum verification successful")
		} else {
			log.Println("md5 checksum verification failed")
		}

		messageHeader, err := response_handlers.GetHeader(conn, decrypted[0:40])
		if err != nil {
			log.Println("failed to read:", err)
			continue
		}
		ReadRemaining := int(messageHeader.MessageLength) - 40
		if ReadRemaining > 0 {
			if messageHeader.TransactionCode == 23001 {
				boxsignonreq, err := response_handlers.ReadBoxSignOnResp(conn, messageHeader, decrypted)
				if err != nil {
					log.Println("failed to read for transcode", messageHeader.TransactionCode, err)
					continue
				}
				log.Print(boxsignonreq)
			} else if messageHeader.TransactionCode == 2301 {
				signonreq, err := response_handlers.ReadSignOnResp(conn, messageHeader, decrypted)
				if err != nil {
					log.Println("failed to read for transcode", messageHeader.TransactionCode, err)
					continue
				}
				log.Print(signonreq)
			} else if messageHeader.TransactionCode == 1601 {

			} else if messageHeader.TransactionCode == 7307 {

			}
		}
	}

}
