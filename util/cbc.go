package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

//CBC 模式

//解密
/**
* rawData 原始加密数据
* key  密钥
* iv  向量
 */
func Decryption(rawData, key, iv string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(rawData)
	key_b, err_1 := base64.StdEncoding.DecodeString(key)
	iv_b, _ := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return "", err
	}
	if err_1 != nil {
		return "", err_1
	}

	dnData, err := AesCBCDecryption(data, key_b, iv_b)
	if err != nil {
		return "", err
	}

	return string(dnData), nil
}

//解密
func AesCBCDecryption(encryptData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	blockSize := block.BlockSize()
	if len(encryptData) < blockSize {
		panic("ciphertext too short")
	}
	if len(encryptData)%blockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encryptData, encryptData)

	//解填充
	encryptData = PKCS7UnPadding(encryptData)

	return encryptData, nil
}

//去除填充
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
