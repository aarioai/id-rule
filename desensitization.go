package id_rule

import (
	"crypto/md5"
	"encoding/hex"
)

const (
	phoneLength = 11
)

// DesensitizeChinesePhoneNumber 手机号脱敏
// 返回脱敏后的手机号、原值的MD5和校验结果
func DesensitizeChinesePhoneNumber(num string) (desensitized string, md5str string, ok bool) {
	if len(num) != phoneLength {
		return "", "", false
	}
	hash := md5.Sum([]byte(num))
	md5str = hex.EncodeToString(hash[:])
	desensitized = num[:3] + "****" + num[7:]
	return desensitized, md5str, true
}

// DesensitizeChineseIdNum 中国身份证号码脱敏
// 返回脱敏后的身份证号、原值的MD5和校验结果
func DesensitizeChineseIdNum(num string) (desensitized string, md5str string, ok bool) {
	if len(num) != idLength {
		return "", "", false
	}
	hash := md5.Sum([]byte(num))
	md5str = hex.EncodeToString(hash[:])
	desensitized = num[:3] + "***" + num[6:14] + "***" + num[17:]
	return desensitized, md5str, true
}
