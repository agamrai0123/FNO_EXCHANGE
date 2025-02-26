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

func SendSignonReq(conn net.Conn, GatewayInfo *models.GatewayRouterResponse, seq uint32) error {
	ReqStructSize := int16(278)
	request := models.SIGN_ON_REQUEST_IN{
		MessageHeader: models.MESSAGE_HEADER{
			TransactionCode: 2300,
			TraderId:        config.TraderId,
			MessageLength:   ReqStructSize,
		},
		UserId:                 config.TraderId,
		Password:               config.Password,
		NewPassword:            config.NewPassword,
		TraderName:             config.TraderName,
		LastPasswordChangeDate: 0,
		BrokerID:               config.BrokerId,
		BranchID:               config.BranchID,
		VersionNumber:          config.VersionNumber,
		Batch2StartTime:        0,
		HostSwitchContext:      0,
		Colour:                 [50]int8{},
		UserType:               config.UserType,
		SequenceNumber:         0.000000,
		WsClassName:            [14]int8{},
		BrokerStatus:           0,
		ShowIndex:              0,
		BrokerEligibilityperMkt: models.ST_BROKER_ELIGIBILITY_PER_MKT{
			MarketEligibilty: 77,
			Reserved1:        0,
		},
		MemberType:     0,
		ClearingStatus: 0,
	}

	// Serialize Data
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, request)
	if err != nil {
		log.Printf("Error while serializing request: %v", err)
		return err
	}
	bufbytes := buf.Bytes()

	// Create MD5 sum
	Checksum := utils.GetMD5Hash(bufbytes)

	// Encrypt data
	Encryptedbuf, err := utils.EncryptAES(GatewayInfo.CryptographicKey, GatewayInfo.CryptographicIV, bufbytes)
	if err != nil {
		log.Printf("Error while writing encryption %v", err)
		return err
	}

	length := make([]byte, 2)
	binary.LittleEndian.PutUint16(length, uint16(request.MessageHeader.MessageLength))
	sequence := make([]byte, 4)
	binary.LittleEndian.PutUint32(sequence, seq)

	packet := append(append(append(length, sequence...), Checksum...), Encryptedbuf...)

	// Write to Socket
	_, err = conn.Write(packet)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return err
	}
	log.Printf("sign-on request sent successfully")
	return nil
}
