package myutils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

// Md5 加密
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// JsonEncode 对数据进行json编码
func JsonEncode[T any](p T) string {
	// 将Person实例编码为JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	// 将JSON字节数组转换为字符串并打印
	return string(jsonData)
}

// JsonDecode json数据解码
func JsonDecode[T any](data []byte, v T) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}
	return nil
}

// GenerateNumericCode 生成指定长度的数字验证码
func GenerateNumericCode(length int) string {
	digits := "0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = digits[rand.Intn(len(digits))]
	}
	return string(result)
}

// EncodeBase64 Base64加密
func EncodeBase64(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// DecodeBase64 Base64解密
func DecodeBase64(encodedData string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}

// Mkdir 创建目录
func Mkdir(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// 目录不存在，创建目录
		err := os.MkdirAll(filePath, 0777)
		if err != nil {
			fmt.Println(222)
			return err
		}
		return nil
	} else if err != nil {
		// 其他错误
		return err
	}
	return nil
}
