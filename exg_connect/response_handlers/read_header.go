package response_handlers

import (
	"encoding/binary"
	"net"

	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/models"
)

func GetHeader(conn net.Conn, buf []byte) (*models.MESSAGE_HEADER, error) {
	var Header models.MESSAGE_HEADER
	// buf := make([]byte, binary.Size(Header))
	// err := utils.SocketRead(conn, buf)
	// if err != nil {
	// 	log.Println("failed to Read:", err)
	// 	return Header, err
	// }
	Header.TransactionCode = int16(binary.LittleEndian.Uint16(buf[0:2]))
	Header.LogTime = int32(binary.LittleEndian.Uint32(buf[2:6]))
	Header.AlphaChar[0] = int8(buf[6])
	Header.AlphaChar[1] = int8(buf[7])
	Header.TraderId = int32(binary.LittleEndian.Uint32(buf[8:12]))
	Header.ErrorCode = int16(binary.LittleEndian.Uint16(buf[12:14]))
	Header.Timestamp = int64(binary.LittleEndian.Uint64(buf[14:22]))
	for i := range 8 {
		Header.TimeStamp1[i] = int8(buf[22+i])
	}
	for i := range 8 {
		Header.TimeStamp2[i] = int8(buf[30+i])
	}
	Header.MessageLength = int16(binary.LittleEndian.Uint16(buf[38:40]))

	return &Header, nil
}
