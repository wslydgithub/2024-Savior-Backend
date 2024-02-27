package tool

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomString(length int) ([]byte, error) {
	randomBytes := make([]byte, length)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	// 将随机字节转换为十六进制字符串
	randomString := hex.EncodeToString(randomBytes)

	// 将字符串转换为 []byte 类型
	randomBytes = []byte(randomString)

	return randomBytes, nil
}
