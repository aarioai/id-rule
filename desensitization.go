package id_rule

import (
	"crypto/md5"
	"encoding/hex"
)

// 数据脱敏隐藏局部，需要多传一个原值的md5用于校验

// 手机号脱敏
func DesensitizeChinesePhoneNumber(num string) (desensitized string, md5str string, ok bool) {
	l := len(num)
	if l < 11 {
		return
	}
	hash := md5.Sum([]byte(num))
	md5str = hex.EncodeToString(hash[:])
	desensitized = num[0:3] + "****" + num[l-4:]
	ok = true
	return
}

// 中国身份证号码脱敏
func DesensitizeChineseIdNum(num string) (desensitized string, md5str string, ok bool) {
	l := len(num)
	if l < 18 {
		return
	}
	hash := md5.Sum([]byte(num))
	md5str = hex.EncodeToString(hash[:])
	desensitized = num[0:3] + "***"+num[6:14]+"***" + num[l-1:]
	ok = true
	return
}
