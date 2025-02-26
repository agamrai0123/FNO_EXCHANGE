package handlers

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/config"
	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/models"
	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/response_handlers"
)

func BoxRegistration(GatewayInfo *models.GatewayRouterResponse) (net.Conn, error) {
	const totalSize = int16(42)
	RespStructSize := int16(40)

	// Create Packet
	request := models.SECURE_BOX_REGISTRATION_REQUEST{
		MessageHeader: models.MESSAGE_HEADER{
			TransactionCode: 23008,
			TraderId:        config.TraderId,
			MessageLength:   totalSize,
		},
		BoxId: GatewayInfo.BoxId,
	}

	// Serialize Data
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, request)
	if err != nil {
		log.Printf("error while serializing request: %v", err)
		return nil, err
	}
	sockInfo := models.SocketInfo{
		Conn_type: "tcp",
		Conn_host: GatewayInfo.IPAddress,
		Conn_port: strconv.FormatInt(int64(GatewayInfo.Port), 10),
		Timeout:   5 * time.Second,
	}
	fmt.Println("GatewayInfo: ", GatewayInfo)
	fmt.Println("SocketInfo: ", sockInfo)

	// Connect to Socket
	conn, err := net.DialTimeout(sockInfo.Conn_type, sockInfo.Conn_host+":"+sockInfo.Conn_port, sockInfo.Timeout)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return nil, err
	}
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		log.Println("Failed to set TCP_NODELAY")
		return nil, err
	}
	if err := tcpConn.SetNoDelay(true); err != nil {
		log.Println("Error setting TCP_NODELAY:", err)
		return nil, err
	}

	// Write to Socket
	_, err = conn.Write(buf.Bytes())
	if err != nil {
		fmt.Println("Error sending message:", err)
		return nil, err
	}
	log.Printf("box registration request sent successfully")

	// Read and parse BOX_REGISTRATION_RESPONSE
	response := make([]byte, RespStructSize)
	_, err = io.ReadFull(conn, response)
	if err != nil {
		fmt.Println("failed to ReadFull:", err)
		return nil, err
	}

	// Read Response Header
	messageHeader, err := response_handlers.GetHeader(conn, response)
	if err != nil {
		log.Printf("error while parsing Box Registration header: %v", err)
		return nil, err
	}
	// Read Response Body
	br := &models.SECURE_BOX_REGISTRATION_RESPONSE{}
	br.MessageHeader = *messageHeader

	log.Printf("box registration response received successfully: %+v", br)
	return conn, nil
}
