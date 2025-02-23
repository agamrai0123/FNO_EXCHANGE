package handler

import (
	"bytes"
	"encoding/binary"
	"log"

	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/config"
	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/models"
	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/utils"
)

func BoxRegistrationHandler(GRResponse *models.GatewayRouterResponse) error {

	StructSize := int16(42)
	request := models.SECURE_BOX_REGISTRATION_REQUEST{
		MessageHeader: models.MESSAGE_HEADER{
			TransactionCode: 23008,
			TraderId:        config.TraderId,
			MessageLength:   StructSize,
		},
		BoxId: GRResponse.BoxId,
	}
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.LittleEndian, request)
	if err != nil {
		log.Printf("Error while serializing request: %v", err)
		return nil, err
	}
	conn, err := utils.SocketConnect(conf, target)
	if err != nil {
		log.Printf("Error while connecting to socket: %v", err)
		return nil, err
	}
	defer conn.Close()
	err = utils.SocketWrite(buf, conn)
	if err != nil {
		log.Printf("Error while writing to socket: %v", err)
		return nil, err
	}
	log.Printf("GR_Request sent successfully")

	// Read and parse GR_RESPONSE
	response := make([]byte, StructSize)
	err = utils.SocketRead(conn, response)
	if err != nil {
		log.Printf("Error while reading from socket: %v", err)
		return nil, err
	}
	grResponse, err := FillGRResponse(response)
	if err != nil {
		log.Printf("Error while parsing GR_RESPONSE header: %v", err)
		return nil, err
	}
	GatewayRouterResponse := ConvertGRToGatewayRouterResponse(grResponse)

	log.Printf("GR_Response received successfully: %+v", grResponse)
	return GatewayRouterResponse, nil
}

// func FillHeader(buf []byte) (*models.MESSAGE_HEADER, error) {
// 	header := &models.MESSAGE_HEADER{}
// 	header.TransactionCode = int16(binary.LittleEndian.Uint16(buf[0:2]))
// 	header.LogTime = int32(binary.LittleEndian.Uint32(buf[2:6]))
// 	header.AlphaChar[0] = int8(buf[6])
// 	header.AlphaChar[1] = int8(buf[7])
// 	header.TraderId = int32(binary.LittleEndian.Uint32(buf[8:12]))
// 	header.ErrorCode = int16(binary.LittleEndian.Uint16(buf[12:14]))
// 	header.Timestamp = int64(binary.LittleEndian.Uint64(buf[14:22]))
// 	for i := 0; i < 8; i++ {
// 		header.TimeStamp1[i] = int8(buf[22+i])
// 		header.TimeStamp2[i] = int8(buf[30+i])
// 	}
// 	header.MessageLength = int16(binary.LittleEndian.Uint16(buf[38:40]))
// 	if header.ErrorCode != 0 {
// 		return header, fmt.Errorf("ErrorCode : %d", header.ErrorCode)
// 	}
// 	return header, nil
// }

// func FillGRResponse(buf []byte) (*models.GR_RESPONSE, error) {
// 	const totalSize = 124
// 	if len(buf) < totalSize {
// 		return nil, fmt.Errorf("buffer too short for GR_RESPONSE: expected %d bytes, got %d", totalSize, len(buf))
// 	}

// 	gr := &models.GR_RESPONSE{}
// 	header, err := FillHeader(buf[0:40])
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to fill header: %w", err)
// 	}
// 	gr.MessageHeader = *header
// 	gr.BoxId = int16(binary.LittleEndian.Uint16(buf[40:42]))
// 	for i := 0; i < 5; i++ {
// 		gr.BrokerId[i] = int8(buf[42+i])
// 	}
// 	gr.Filler = int8(buf[47])
// 	for i := 0; i < 16; i++ {
// 		gr.IPAddress[i] = int8(buf[48+i])
// 	}
// 	gr.Port = int32(binary.LittleEndian.Uint32(buf[64:68]))
// 	for i := 0; i < 8; i++ {
// 		gr.SessionKey[i] = int8(buf[68+i])
// 	}
// 	for i := 0; i < 32; i++ {
// 		gr.CryptographicKey[i] = int8(buf[76+i])
// 	}
// 	for i := 0; i < 16; i++ {
// 		gr.CryptographicIV[i] = int8(buf[108+i])
// 	}

// 	return gr, nil
// }

// func ConvertGRToGatewayRouterResponse(gr *models.GR_RESPONSE) *models.GatewayRouterResponse {
// 	return &models.GatewayRouterResponse{
// 		IPAddress:        utils.Int8SliceToString(gr.IPAddress[:]),
// 		Port:             gr.Port,
// 		SessionKey:       utils.Int8SliceToString(gr.SessionKey[:]),
// 		CryptographicKey: utils.Int8SliceToString(gr.CryptographicKey[:]),
// 		CryptographicIV:  utils.Int8SliceToString(gr.CryptographicIV[:]),
// 	}
// }
