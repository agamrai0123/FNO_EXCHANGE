package response_handlers

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"net"

	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/models"
)

func ReadSignOnResp(conn net.Conn, Header *models.MESSAGE_HEADER, buf []byte) (*models.SIGN_ON_REQUEST_OUT, error) {
	const totalSize = 278
	signOnResp := &models.SIGN_ON_REQUEST_OUT{}
	signOnResp.MessageHeader = *Header
	ReadRemaining := int(Header.MessageLength) - 40
	if ReadRemaining < (totalSize - 40) {
		return nil, fmt.Errorf("buffer too short for GR_RESPONSE: expected %d bytes, got %d", totalSize, ReadRemaining)
	}
	if ReadRemaining <= 0 {
		return nil, errors.New("invalid message length")
	}
	signOnResp.UserId = int32(binary.LittleEndian.Uint32(buf[40:44]))
	for i := range 8 {
		signOnResp.Password[i] = int8(buf[52+i])
		signOnResp.NewPassword[i] = int8(buf[68+i])
	}
	for i := range 26 {
		signOnResp.TraderName[i] = int8(buf[76+i])
	}
	signOnResp.LastPasswordChangeDate = int32(binary.LittleEndian.Uint32(buf[102:106]))
	for i := range 5 {
		signOnResp.BrokerID[i] = int8(buf[106+i])
	}
	signOnResp.BranchID = int16(binary.LittleEndian.Uint16(buf[112:114]))
	signOnResp.VersionNumber = int32(binary.LittleEndian.Uint32(buf[114:118]))
	signOnResp.EndTime = int32(binary.LittleEndian.Uint32(buf[118:122]))
	for i := range 50 {
		signOnResp.Colour[i] = int8(buf[123+i])
	}
	signOnResp.UserType = int16(binary.LittleEndian.Uint16(buf[174:176]))
	signOnResp.SequenceNumber = math.Float64frombits(binary.LittleEndian.Uint64(buf[176:184]))
	signOnResp.BrokerStatus = int8(buf[198])
	signOnResp.ShowIndex = int8(buf[199])
	signOnResp.BrokerEligibilityPerMkt.MarketEligibilty = uint8(buf[200])
	signOnResp.MemberType = int16(binary.LittleEndian.Uint16(buf[202:204]))
	signOnResp.ClearingStatus = int8(buf[204])
	for i := range 25 {
		signOnResp.BrokerName[i] = int8(buf[205+i])
	}
	return signOnResp, nil
}
