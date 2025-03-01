package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func EncryptAES(Key, InitVector, plainText []byte) ([]byte, error) {
	block, err := aes.NewCipher(Key)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM mode: %w", err)
	}
	nonce := InitVector
	if len(InitVector) != gcm.NonceSize() {
		nonce = InitVector[:gcm.NonceSize()]
	}
	ciphertext := gcm.Seal(nil, nonce, plainText, nil)
	return ciphertext, nil
}

func DecryptAES(Key, InitVector, cyphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(Key)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM mode: %w", err)
	}
	nonce := InitVector
	if len(InitVector) != gcm.NonceSize() {
		nonce = InitVector[:gcm.NonceSize()]
	}
	plaintext, err := gcm.Open(nil, nonce, cyphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt: %w", err)
	}
	return plaintext, nil
}
