package id_rule

import (
	"math"
	"strconv"
	"time"
)

// The Rule of Chinese Citizen Identification Number 中国身份证ID校验
// 计算规则参考“中国国家标准化管理委员会”官方文档：http://www.gb688.cn/bzgk/gb/newGbInfo?hcno=080D6FBF2BB468F9007657F26D60013E

type ChineseCitizenId struct {
	valid      bool
	districtId uint64
	birthDate  string
	sex        Sex
}

func NewChineseCitizenId(num string) ChineseCitizenId {
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

// 检查是否符合身份证国标
func isChineseCitizenIdValid(num string) bool {
	//a1与对应的校验码对照表，其中key表示a1，value表示校验码，value中的10表示校验码X
	var a1Map = map[int]int{
		0:  1,
		1:  0,
		2:  10,
		3:  9,
		4:  8,
		5:  7,
		6:  6,
		7:  5,
		8:  4,
		9:  3,
		10: 2,
	}

	var sum int
	var signChar = ""
	for index, c := range num {
		var i = 18 - index
		if i != 1 {
			v, err := strconv.Atoi(string(c))
			if err != nil {
				return false
			}
			//计算加权因子
			var weight = int(math.Pow(2, float64(i-1))) % 11
			sum += v * weight

		} else {
			signChar = string(c)
		}
	}
	var a1 = a1Map[sum%11]
	var a1Str string
	if a1 == 10 {
		a1Str = "X"
	} else {
		a1Str = strconv.Itoa(a1)
	}
	return a1Str == signChar
}

// 解析中国身份证号码，原籍（省份、城市）、生日、性别, 是否正确（通过身份证自带简易校验）
func parseChineseCitizenId(num string) (district uint64, birthDate string, sex Sex, valid bool) {
	var (
		x   int
		err error
	)

	birthDate = num[6:10] + "-" + num[10:12] + "-" + num[12:14]
	_, err = time.Parse("2006-01-02", birthDate)
	if err != nil {
		return
	}

	veri := num[14:17]
	x, err = strconv.Atoi(veri)
	if err != nil {
		return
	}

	sex = UnknownSex
	if x%2 == 0 {
		sex = Female
	} else {
		sex = Male
	}
	district, err = strconv.ParseUint(num[0:6], 10, 64)
	if err != nil {
		return
	}

	valid = isChineseCitizenIdValid(num)
	if !valid {
		return
	}

	valid = true
	return
}
