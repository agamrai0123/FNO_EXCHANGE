package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"io"
	"net"

	"github.com/agamrai0123/FNO_EXCHANGE/exg_connect/models"
)

// func Int8SliceToString(arr []int8) string {
// 	b := make([]byte, len(arr))
// 	for i, v := range arr {
// 		b[i] = byte(v)
// 	}
// 	b = bytes.Trim(b, "\x00")
// 	return string(b)
// }

// func StringtoInt8Slice(str string) []int8 {
// 	byteArr := []byte(str)
// 	int8Arr := make([]int8, len(byteArr))
// 	for i, b := range byteArr {
// 		int8Arr[i] = int8(b)
// 	}
// 	return int8Arr
// }

func GetMD5Hash(data []byte) []byte {
	hash := md5.Sum(data)
	// return hex.EncodeToString(hash[:])
	return hash[:]
}

func EncryptAES(Key, Vector, plainText []byte) ([]byte, error) {
	block, err := aes.NewCipher(Key)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM mode: %w", err)
	}
	ciphertext := gcm.Seal(nil, Vector, plainText, nil)
	return ciphertext, nil
}

func DecryptAES(Key, Vector, cyphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(Key)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM mode: %w", err)
	}
	plaintext, err := gcm.Open(nil, Vector, cyphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt: %w", err)
	}
	return plaintext, nil
}

func ReadExchangePacket(conn net.Conn) (*models.ExchangeData, error) {
	buf := make([]byte, 1024)
	_, err := io.ReadFull(conn, buf)
	if err != nil {
		fmt.Println("failed to ReadFull:", err)
		return nil, err
	}
	exgdata := models.ExchangeData{
		Length:         int16(binary.LittleEndian.Uint16(buf[0:2])),
		SequenceNumber: int32(binary.LittleEndian.Uint32(buf[2:6])),
		Checksum:       buf[6:22],
		MessageData:    buf[22:],
	}
	return &exgdata, nil
}
