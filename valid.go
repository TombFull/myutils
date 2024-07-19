package myutils

import "regexp"

// IsValidMobile 验证是否是正确的手机号
func IsValidMobile(mobile string) bool {
	// 使用正则表达式验证手机号格式（以中国大陆手机号为例）
	regex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	return regex.MatchString(mobile)
}
