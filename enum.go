package idcheck

type Sex uint8

const (
	UnknownSex Sex = 0
	Male       Sex = 1
	Female     Sex = 2
	OtherSex   Sex = 255
)
