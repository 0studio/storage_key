package key

import (
	"strconv"
	"strings"
)

type KeyInt32 int32
type Int32 KeyInt32

func (u KeyInt32) Equal(u2 KeyInt32) bool {
	return u.ToInt32() == u2.ToInt32()
}

func (u *KeyInt32) FromString(str string) bool {
	value, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	*u = KeyInt32(value)
	return true
}

func (u KeyInt32) String() string {
	return u.ToString()
}
func (u KeyInt32) ToString() string {
	return strconv.Itoa(int(u))
}
func (u KeyInt32) ToInt32() int32 {
	return int32(u)
}

func (u KeyInt32) ToInt() int {
	return int(u)
}
func (u KeyInt32) ToSum() int {
	return int(u)
}
func (u KeyInt32) SumLen() int {
	return 1
}
func (u KeyInt32) GetSumByIdx(_ int) Sum {
	return u
}

type KeyInt32List []KeyInt32

func (keys KeyInt32List) SumLen() int {
	return len(keys)
}
func (keys KeyInt32List) ToString() string {
	return keys.Join("|")
}
func (keys KeyInt32List) ToSum() (sum int) {
	for _, k := range keys {
		sum += k.ToSum()
	}
	return
}
func (keys KeyInt32List) GetSumByIdx(idx int) Sum {
	if idx < len(keys) && idx > -1 {
		return keys[idx]
	}
	return nil
}

func (keys KeyInt32List) ToStringList() (strList []string) {
	strList = make([]string, len(keys))
	for idx, _ := range keys {
		strList[idx] = keys[idx].ToString()
	}
	return
}
func (keys KeyInt32List) Join(sep string) string {
	return strings.Join(keys.ToStringList(), sep)
}

func (keys KeyInt32List) IsInList(Key KeyInt32) bool {
	for _, tmpKey := range keys {
		if tmpKey.Equal(Key) {
			return true
		}
	}
	return false
}

func (keys KeyInt32List) Remove(Key KeyInt32) (ret KeyInt32List) {
	ret = make(KeyInt32List, len(keys), len(keys))
	var index int = 0
	for _, value := range keys {
		if !value.Equal(Key) {
			ret[index] = value
			index++
		}
	}
	ret = ret[:index]
	return

}
