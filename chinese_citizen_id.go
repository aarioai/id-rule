package id_rule

import (
	"strconv"
	"time"
)

// The Rule of Chinese Citizen Identification Number 中国身份证ID校验
// 计算规则参考“中国国家标准化管理委员会”官方文档：http://www.gb688.cn/bzgk/gb/newGbInfo?hcno=080D6FBF2BB468F9007657F26D60013E

const (
	idLength      = 18
	districtLen   = 6
	birthDateLen  = 8
	sequenceLen   = 3
	checksumIndex = 17
	birthStartPos = 6
	birthEndPos   = 14
)

type ChineseCitizenId struct {
	valid      bool
	districtId uint64
	birthDate  string
	sex        Sex
}

// NewChineseCitizenId 创建新的身份证验证实例
// 如果输入的身份证号码格式不正确，将返回一个无效的实例
func NewChineseCitizenId(num string) ChineseCitizenId {
	if len(num) != idLength {
		return ChineseCitizenId{valid: false}
	}
	d, b, s, v := parseChineseCitizenId(num)
	return ChineseCitizenId{
		valid:      v,
		districtId: d,
		birthDate:  b,
		sex:        s,
	}
}

func (n ChineseCitizenId) Valid() bool {
	return n.valid
}
func (n ChineseCitizenId) DistrictId() uint64 {
	return n.districtId
}
func (n ChineseCitizenId) BirthDate() string {
	return n.birthDate
}
func (n ChineseCitizenId) Sex() Sex {
	return n.sex
}

// isChineseCitizenIdValid 优化后的身份证校验算法
func isChineseCitizenIdValid(num string) bool {
	// 预计算权重因子，避免重复计算
	var weights = [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var validChecksum = [11]string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}

	var sum int
	for i := 0; i < len(weights); i++ {
		digit, err := strconv.Atoi(string(num[i]))
		if err != nil {
			return false
		}
		sum += digit * weights[i]
	}

	checksum := validChecksum[sum%11]
	return string(num[checksumIndex]) == checksum
}

func parseChineseCitizenId(num string) (district uint64, birthDate string, sex Sex, valid bool) {
	// 首先验证长度
	if len(num) != idLength {
		return 0, "", UnknownSex, false
	}

	// 验证出生日期
	birthDate = num[birthStartPos:birthStartPos+4] + "-" +
		num[birthStartPos+4:birthStartPos+6] + "-" +
		num[birthStartPos+6:birthEndPos]

	if t, err := time.Parse("2006-01-02", birthDate); err != nil || t.After(time.Now()) {
		return 0, "", UnknownSex, false
	}

	// 验证地区码
	district, err := strconv.ParseUint(num[0:districtLen], 10, 64)
	if err != nil {
		return 0, "", UnknownSex, false
	}

	// 解析性别
	if sequence, err := strconv.Atoi(num[14:17]); err == nil {
		if sequence%2 == 0 {
			sex = Female
		} else {
			sex = Male
		}
	} else {
		return 0, "", UnknownSex, false
	}

	// 验证校验码
	valid = isChineseCitizenIdValid(num)
	return
}
