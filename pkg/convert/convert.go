package convert

import "strconv"

type strTo string

func (s strTo) String() string {
	return string(s)
}

func (s strTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

func (s strTo) MustInt() int {
	v, _ := s.Int()
	return v
}

func (s strTo) UInt() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

func (s strTo) MustUInt32() uint32 {
	v, _ := s.UInt()
	return v
}