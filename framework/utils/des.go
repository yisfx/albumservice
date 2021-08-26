package utils

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"fmt"
)

func DesDemo() error {
	data := []byte("hello world")
	key := []byte("12345678")
	iv := []byte("43218765")
	result, err := DESCBCEncrypt(data, key, iv)
	if err != nil {
		return err
	}
	b := hex.EncodeToString(result)
	fmt.Println("en:", b)
	//-----------------------------------
	if err != nil {
		return err
	}
	result, err = DESCBCDecrypt(result, key, iv)
	if err != nil {
		return err
	}
	fmt.Println("de:", string(result))
	return nil
}

func DESCBCEncrypt(data, key, iv []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	data = pkcs5Padding(data, block.BlockSize())
	cryptText := make([]byte, len(data))

	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(cryptText, data)
	return cryptText, nil
}
func DESCBCDecrypt(data, key, iv []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	cryptText := make([]byte, len(data))
	blockMode.CryptBlocks(cryptText, data)
	cryptText = pkcs5UnPadding(cryptText)
	return cryptText, nil
}

func pkcs5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:length-unpadding]
}
