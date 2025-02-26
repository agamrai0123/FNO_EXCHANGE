package request_handlers

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"

	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/config"
	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/models"
	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/utils"
)

func SendBoxSignonReq(conn net.Conn, GatewayInfo *models.GatewayRouterResponse, seq uint32) error {
	ReqStructSize := int16(60)
	// Create Packet
	request := models.BOX_SIGN_ON_REQUEST_IN{
		MessageHeader: models.MESSAGE_HEADER{
			TransactionCode: 23000,
			TraderId:        config.TraderId,
			MessageLength:   ReqStructSize,
		},
		BoxId:      GatewayInfo.BoxId,
		BrokerId:   config.BrokerId,
		SessionKey: GatewayInfo.SessionKey,
	}

	// Serialize Data
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, request)
	if err != nil {
		log.Printf("error while serializing request: %v", err)
		return err
	}

	bufbytes := buf.Bytes()

	// Create MD5 sum
	checksum := utils.GetMD5Hash(bufbytes)

	// Encrypt data
	encryptedBuf, err := utils.EncryptAES(GatewayInfo.CryptographicKey, GatewayInfo.CryptographicIV, bufbytes)
	if err != nil {
		log.Printf("error while writing encryption %v", err)
		return err
	}

	length := make([]byte, 2)
	binary.LittleEndian.PutUint16(length, uint16(request.MessageHeader.MessageLength))
	sequence := make([]byte, 4)
	binary.LittleEndian.PutUint32(sequence, seq)
	packet := append(append(append(length, sequence...), checksum...), encryptedBuf...)

	// Write to Socket
	_, err = conn.Write(packet)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return err
	}
	log.Printf("box sign-on request sent successfully")
	return nil
}
