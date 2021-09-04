package utils

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"fmt"
)

var DESHelper deshelper

type deshelper struct {
	key []byte
	iv  []byte
}

func init() {
	DESHelper = deshelper{}
}

func DesDemo(data string) error {
	result, err := DESHelper.DESCBCEncrypt(data)
	if err != nil {
		return err
	}
	fmt.Println("en:", result)
	//-----------------------------------
	if err != nil {
		return err
	}
	result, err = DESHelper.DESCBCDecrypt(result)
	if err != nil {
		return err
	}
	fmt.Println("de:", string(result))
	return nil
}

func (d *deshelper) Init(k, i string) {
	d.key = []byte(k)
	d.iv = []byte(i)
}

func (d *deshelper) DESCBCEncrypt(str string) (string, error) {
	data := []byte(str)
	block, err := des.NewCipher(d.key)
	if err != nil {
		return "", err
	}
	data = pkcs5Padding(data, block.BlockSize())
	cryptText := make([]byte, len(data))

	blockMode := cipher.NewCBCEncrypter(block, d.iv)
	blockMode.CryptBlocks(cryptText, data)
	result := hex.EncodeToString(cryptText)
	return result, nil
}

func (d *deshelper) DESCBCDecrypt(str string) (string, error) {
	data, _ := hex.DecodeString(str)
	block, err := des.NewCipher(d.key)
	if err != nil {
		return "", err
	}

	blockMode := cipher.NewCBCDecrypter(block, d.iv)
	cryptText := make([]byte, len(data))
	blockMode.CryptBlocks(cryptText, data)
	cryptText = pkcs5UnPadding(cryptText)
	return string(cryptText), nil
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
