package id_rule

import (
	"math"
	"regexp"
	"strings"
)

const (
	usciLength = 18
	base       = 31
)

var valueMap = map[int32]int{
	'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
	'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14, 'F': 15, 'G': 16, 'H': 17, 'J': 18,
	'K': 19, 'L': 20, 'M': 21, 'N': 22, 'P': 23, 'Q': 24, 'R': 25, 'T': 26, 'U': 27,
	'W': 28, 'X': 29, 'Y': 30,
}

type ChineseOrgId string

// NewChineseOrgId 创建新的组织机构代码实例
func NewChineseOrgId(usci string) ChineseOrgId {
	return ChineseOrgId(usci)
}

// IsValid 验证组织机构代码的有效性
func (usci ChineseOrgId) IsValid() bool {
	usciStr := strings.ToUpper(string(usci))
	if len(usciStr) != usciLength {
		return false
	}

	reg, err := regexp.Compile(`^[A-Z0-9]{18}$`)
	if err != nil || !reg.MatchString(usciStr) {
		return false
	}

	sum := 0
	for index, c := range usciStr[:17] {
		value, exists := valueMap[c]
		if !exists {
			return false
		}
		weight := int(math.Pow(3, float64(index))) % base
		sum += value * weight
	}

	mod := sum % base
	sign := base - mod
	if sign == base {
		sign = 0
	}

	signChar := getSignChar(sign)
	return string(usciStr[17]) == signChar
}

// getSignChar 根据计算的校验值获取对应的字符
func getSignChar(sign int) string {
	for key, value := range valueMap {
		if value == sign {
			return string(key)
		}
	}
	return ""
}
