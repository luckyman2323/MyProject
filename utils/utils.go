package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

type EncodeType int

const (
	BASE64 EncodeType = iota + 1
	HEX
)

var AESKeys = []byte("www.yunphant.com")

// aes加密, 支持base64/hex编码
func AesEncrypt(origData []byte, encodeType EncodeType) (string, error) {
	block, err := aes.NewCipher(AESKeys)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	fmt.Println("blockSize: ", blockSize)
	//src must be a multiple of the block size
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, AESKeys[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)

	var res string
	switch encodeType {
	case BASE64:
		res = base64.StdEncoding.EncodeToString(crypted)
	case HEX:
		res = hex.EncodeToString(crypted)
	}
	return res, nil
}

// aes解密, 支持base64/hex编码
func AesDecrypt(origData string, encodeType EncodeType) (string, error) {
	var crypted []byte
	var err error
	switch encodeType {
	case BASE64:
		crypted, err = base64.StdEncoding.DecodeString(origData)
		if err != nil {
			return "", err
		}
	case HEX:
		crypted, err = hex.DecodeString(origData)
		if err != nil {
			return "", err
		}
	}

	block, err := aes.NewCipher(AESKeys)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, AESKeys[:blockSize])
	res := make([]byte, len(crypted))
	blockMode.CryptBlocks(res, crypted)
	res = PKCS5UnPadding(res)
	return string(res), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	if length == 0 {
		panic("origData is empty")
	}
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
