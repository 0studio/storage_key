package key

import (
	"strconv"
)

type KeyInt int

func (u KeyInt) Equal(u2 KeyInt) bool {
	return u.ToUint64() == u2.ToUint64()
}

func (u *KeyInt) FromString(str string) bool {
	value, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	*u = KeyInt(value)
	return true
}

func (u KeyInt) ToString() string {
	return strconv.Itoa(u.ToInt())
}
func (u KeyInt) ToInt() int {
	return int(u)
}
func (u KeyInt) ToUint64() uint64 {
	return uint64(u)
}
func (u KeyInt) ToUnt64() int64 {
	return int64(u)
}
