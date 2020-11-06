package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/skip2/go-qrcode"
	"github.com/spf13/viper"
	"log"
)

type QrInfo struct {
	UserId int    `json:"user_id"`
	Source string `json:"source"`
}

//RSA加密生成二维码
func CreateQrCode(data string) ([]byte, error) {
	publicKey := viper.GetString("qrCode.publicKey")
	result, err := RsaEncryptWithSha1Base64(data, publicKey)
	if err != nil {
		return nil, err
	}

	var png []byte
	png, err = qrcode.Encode(result, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}

	return png, nil
}

//解析二维码
func ParseQrCode(param string, codeInfo interface{}) (bool, error) {
	privateKey := viper.GetString("qrCode.privateKey")
	result, err := RsaDecryptWithSha1Base64(param, privateKey)
	if err != nil {
		return false, err
	}

	//解二维码
	log.Println("result is ", result)
	err = json.Unmarshal([]byte(result), &codeInfo)
	fmt.Printf("codeInfo is %v", codeInfo)
	if err != nil {
		return false, err
	}

	return true, nil
}

func RsaEncryptWithSha1Base64(originalData, publicKey string) (string, error) {
	key, _ := base64.StdEncoding.DecodeString(publicKey)
	pubKey, _ := x509.ParsePKIXPublicKey(key)
	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey.(*rsa.PublicKey), []byte(originalData))

	return base64.StdEncoding.EncodeToString(encryptedData), err
}

func RsaDecryptWithSha1Base64(encryptedData, privateKey string) (string, error) {
	encryptedDecodeBytes, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}

	key, _ := base64.StdEncoding.DecodeString(privateKey)
	prvKey, err := x509.ParsePKCS1PrivateKey(key)
	originalData, err := rsa.DecryptPKCS1v15(rand.Reader, prvKey, encryptedDecodeBytes)

	return string(originalData), err
}
