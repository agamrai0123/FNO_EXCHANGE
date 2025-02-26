package response_handlers

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"

	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/models"
)

func ReadBoxSignOnResp(conn net.Conn, Header *models.MESSAGE_HEADER, buf []byte) (*models.BOX_SIGN_ON_REQUEST_OUT, error) {
	const totalSize = 52
	bsr := &models.BOX_SIGN_ON_REQUEST_OUT{}
	bsr.MessageHeader = *Header
	ReadRemaining := int(Header.MessageLength) - 40
	if ReadRemaining < (totalSize - 40) {
		return nil, fmt.Errorf("buffer too short for GR_RESPONSE: expected %d bytes, got %d", totalSize, ReadRemaining)
	}
	if ReadRemaining <= 0 {
		return nil, errors.New("invalid message length")
	}
	bsr.BoxId = int16(binary.LittleEndian.Uint16(buf[40:42]))
	for i := range 5 {
		bsr.Reserved1[i] = int8(buf[42+i])
	}
	return bsr, nil
}
