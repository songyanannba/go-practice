package helper

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"errors"
)

type AES struct {
	Key []byte
	Iv  []byte
}

func NewAES(key []byte, iv ...[]byte) *AES {
	var iv_ []byte
	if len(iv) > 0 && len(iv[0]) > 0 {
		iv_ = iv[0]
	} else {
		iv_ = key
	}
	return &AES{Key: key, Iv: iv_}
}

// Encrypt
func (a *AES) Encrypt(origData []byte) ([]byte, error) {
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = pKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, a.Iv[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	// hex encode
	encode := hex.EncodeToString(crypted)
	// Add head and tail 8 letter
	//encode = RandomString(9) + encode + RandomString(9)
	return []byte(encode), nil
}

// Decrypt
func (a *AES) Decrypt(crypted []byte) ([]byte, error) {
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return nil, err
	}
	// Trim head and tail 8 letter
	//crypted = crypted[9 : len(crypted)-9]
	// hex decode
	if crypted, err = hex.DecodeString(string(crypted)); err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, a.Iv[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData, err = pKCS7UnPadding(origData)
	return origData, err
}

func pKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pKCS7UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	unpadding := int(origData[length-1])
	if length < unpadding {
		return origData, errors.New("slice bounds out of range")
	}
	return origData[:(length - unpadding)], nil
}

func AesGcmEncrypt(key, nonce, plaintext []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesgcm, err := cipher.NewGCMWithNonceSize(block, 24)
	if err != nil {
		return "", err
	}
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func AesGcmDecrypt(key, nonce, ciphertext []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesgcm, err := cipher.NewGCMWithNonceSize(block, 24)
	if err != nil {
		return "", err
	}
	s, err := base64.StdEncoding.DecodeString(string(ciphertext))
	if err != nil {
		return "", err
	}
	plaintext, err := aesgcm.Open(nil, nonce, s, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
