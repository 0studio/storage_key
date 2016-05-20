package key

import (
	"fmt"
	"strconv"
)

type KeyDoubleUint64 struct {
	key1 uint64
	key2 uint64
}

func NewKeyDoubleUint64(key1, key2 uint64) KeyDoubleUint64 {
	return KeyDoubleUint64{key1: key1, key2: key2}
}
func (u KeyDoubleUint64) Equal(u2 KeyDoubleUint64) bool {
	return u.key1 == u2.key1 && u.key2 == u2.key2
}

func (u *KeyDoubleUint64) FromString(str string) bool {
	_, err := fmt.Sscanf(str, "%d-%d", &u.key1, &u.key2)
	return err == nil
}

func (u KeyDoubleUint64) String() string {
	return u.ToString()
}
func (u KeyDoubleUint64) ToString() string {
	return strconv.FormatUint(u.key1, 10) + "-" + strconv.FormatUint(u.key2, 10)
}

func (u KeyDoubleUint64) ToInt() int {
	return int(u.key1 + u.key2)
}
func (u KeyDoubleUint64) SumLen() int {
	return 1
}
func (keys KeyDoubleUint64) ToSum() (sum int) {
	return int(keys.key1 + keys.key2)
}

func (u KeyDoubleUint64) GetSumByIdx(_ int) Sum {
	return u
}

func (u KeyDoubleUint64) GetKey1() (key1 uint64) {
	key1 = u.key1
	return
}
func (u KeyDoubleUint64) GetKey2() (key2 uint64) {
	key2 = u.key2
	return
}

type KeyDoubleUint64List []KeyDoubleUint64

func (keys KeyDoubleUint64List) Remove(Key KeyDoubleUint64) (ret KeyDoubleUint64List) {
	ret = make(KeyDoubleUint64List, len(keys), len(keys))
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

func (keys KeyDoubleUint64List) IsInList(Key KeyDoubleUint64) bool {
	for _, tmpKey := range keys {
		if tmpKey.Equal(Key) {
			return true
		}
	}
	return false
}
func (u KeyDoubleUint64List) ToInt() int {
	return 0
}
func (u KeyDoubleUint64List) SumLen() int {
	return len(u)
}
func (u KeyDoubleUint64List) GetSumByIdx(idx int) Sum {
	return u[idx]
}
