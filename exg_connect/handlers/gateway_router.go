package handlers

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/config"
	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/models"
	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/response_handlers"
	"github.com/spf13/viper"
)

func GatewayRouter() (net.Conn, *models.GatewayRouterResponse, error) {
	target := "box"
	const totalSize = int16(48)

	// Load configuration
	conf, err := loadConfig()
	if err != nil {
		log.Println("Error during loading config")
		return nil, nil, err
	}

	// Create Packet
	request := &models.GR_REQUEST{
		MessageHeader: models.MESSAGE_HEADER{
			TransactionCode: 2400,
			TraderId:        config.TraderId,
			MessageLength:   totalSize,
		},
		BoxId:    int16(conf.GetInt(target + ".BOX_ID")),
		BrokerId: config.BrokerId,
	}

	// Serialize Data
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.LittleEndian, request)
	if err != nil {
		log.Printf("Error while serializing request: %v", err)
		return nil, nil, err
	}

	// Connect to Socket
	conn, err := net.DialTimeout(conf.GetString(target+".CONN_TYPE"),
		conf.GetString(target+".CONN_HOST")+":"+conf.GetString(target+".CONN_PORT"),
		5*time.Second)
	if err != nil {
		log.Println("Error connecting to server:", err)
		return nil, nil, err
	}
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		log.Println("Failed to set TCP_NODELAY")
		return nil, nil, err
	}
	if err := tcpConn.SetNoDelay(true); err != nil {
		log.Println("Error setting TCP_NODELAY:", err)
		return nil, nil, err
	}

	// Write to Socket
	_, err = conn.Write(buf.Bytes())
	if err != nil {
		log.Println("Error sending message:", err)
		return nil, nil, err
	}
	log.Printf("gateway router request sent successfully")

	// Read and parse Gateway Router Response
	// Read Header
	readbuf := make([]byte, 40)
	_, err = io.ReadFull(conn, readbuf)
	if err != nil {
		log.Println("failed to ReadFull:", err)
		return nil, nil, err
	}

	messageHeader, err := response_handlers.GetHeader(conn, readbuf)
	if err != nil {
		if err.Error() == "EOF" {
			log.Println("nothing to read")
		} else {
			log.Println("failed to read")
		}
	}

	// Read Response Body
	gatewayResp, err := getGatewayResp(conn, messageHeader)
	if err != nil {
		log.Println("failed to read grResponse", messageHeader.TransactionCode, err)
	}
	log.Printf("GR_Response received successfully: %+v", gatewayResp)

	gatewayRouterResp := parseGatewayRouterResp(gatewayResp)
	return conn, gatewayRouterResp, nil
}

func getGatewayResp(conn net.Conn, Header *models.MESSAGE_HEADER) (*models.GR_RESPONSE, error) {
	const totalSize = 124
	gatewayRouterResp := &models.GR_RESPONSE{}
	gatewayRouterResp.MessageHeader = *Header
	ReadRemaining := int(Header.MessageLength) - 40
	if ReadRemaining < (totalSize - 40) {
		return nil, fmt.Errorf("buffer too short for GR_RESPONSE: expected %d bytes, got %d", totalSize, ReadRemaining)
	}
	if ReadRemaining <= 0 {
		return nil, errors.New("invalid message length")
	}

	buf := make([]byte, ReadRemaining)
	_, err := io.ReadFull(conn, buf)
	if err != nil {
		log.Println("failed to ReadFull:", err)
		return nil, err
	}
	gatewayRouterResp.BoxId = int16(binary.LittleEndian.Uint16(buf[0:2]))
	for i := range 5 {
		gatewayRouterResp.BrokerId[i] = int8(buf[2+i])
	}
	gatewayRouterResp.Filler = int8(buf[7])
	for i := range 16 {
		gatewayRouterResp.IPAddress[i] = int8(buf[8+i])
	}
	gatewayRouterResp.Port = int32(binary.LittleEndian.Uint32(buf[24:28]))
	for i := range 8 {
		gatewayRouterResp.SessionKey[i] = int8(buf[28+i])
	}
	for i := range 32 {
		gatewayRouterResp.CryptographicKey[i] = int8(buf[36+i])
	}
	for i := range 16 {
		gatewayRouterResp.CryptographicIV[i] = int8(buf[68+i])
	}
	return gatewayRouterResp, nil
}

func parseGatewayRouterResp(gr *models.GR_RESPONSE) *models.GatewayRouterResponse {
	return &models.GatewayRouterResponse{
		IPAddress:        Int8SliceToString(gr.IPAddress[:]),
		Port:             gr.Port,
		BoxId:            gr.BoxId,
		SessionKey:       gr.SessionKey,
		CryptographicKey: parseKeyBytes(gr.CryptographicKey),
		CryptographicIV:  parseIVBytes(gr.CryptographicIV),
	}
}

func parseKeyBytes(key [32]int8) []byte {
	b := make([]byte, 32)
	for i := range 32 {
		b[i] = byte(key[i])
	}
	return b
}

func parseIVBytes(vector [16]int8) []byte {
	b := make([]byte, 16)
	for i := range 16 {
		b[i] = byte(vector[i])
	}
	return b
}

func loadConfig() (*viper.Viper, error) {
	config := viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName("config")
	config.AddConfigPath("./../config")
	if err := config.ReadInConfig(); err != nil {
		log.Fatal("error on parsing configuration file", err)
		return nil, err
	}
	log.Println("config loaded successfully")
	return config, nil
}

func Int8SliceToString(arr []int8) string {
	b := make([]byte, len(arr))
	for i, v := range arr {
		b[i] = byte(v)
	}
	b = bytes.Trim(b, "\x00")
	return string(b)
}
