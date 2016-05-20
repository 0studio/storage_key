package key

import (
	"fmt"
	"strconv"
)

type KeyUint64Int32 struct {
	key1 uint64
	key2 int32
}

func NewKeyUint64Int32Key(key1 KeyUint64, key2 KeyInt32) KeyUint64Int32 {
	return KeyUint64Int32{key1: key1.ToUint64(), key2: key2.ToInt32()}
}
func NewKeyUint64Int32(key1 uint64, key2 int32) KeyUint64Int32 {
	return KeyUint64Int32{key1: key1, key2: key2}
}
func (u KeyUint64Int32) Equal(u2 KeyUint64Int32) bool {
	return u.key1 == u2.key1 && u.key2 == u2.key2
}

func (u *KeyUint64Int32) FromString(str string) bool {
	_, err := fmt.Sscanf(str, "%d-%d", &u.key1, &u.key2)
	return err == nil
}

func (u KeyUint64Int32) String() string {
	return strconv.FormatUint(u.key1, 10) + "-" + strconv.Itoa(int(u.key2))
}
func (u KeyUint64Int32) ToString() string {
	return u.String()
}

func (u KeyUint64Int32) ToInt() int {
	return int(u.key1) + int(u.key2)
}
func (u KeyUint64Int32) SumLen() int {
	return 1
}
func (keys KeyUint64Int32) ToSum() (sum int) {
	return keys.ToInt()
}

func (u KeyUint64Int32) GetSumByIdx(_ int) Sum {
	return u
}

func (u KeyUint64Int32) GetKey1() (key1 uint64) {
	key1 = u.key1
	return
}
func (u KeyUint64Int32) GetKey2() (key2 int32) {
	key2 = u.key2
	return
}

type KeyUint64Int32List []KeyUint64Int32

func (keys KeyUint64Int32List) Remove(Key KeyUint64Int32) (ret KeyUint64Int32List) {
	ret = make(KeyUint64Int32List, len(keys), len(keys))
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

func (keys KeyUint64Int32List) IsInList(Key KeyUint64Int32) bool {
	for _, tmpKey := range keys {
		if tmpKey.Equal(Key) {
			return true
		}
	}
	return false
}
func (u KeyUint64Int32List) ToInt() int {
	return 0
}
func (u KeyUint64Int32List) SumLen() int {
	return len(u)
}
func (u KeyUint64Int32List) GetSumByIdx(idx int) Sum {
	return u[idx]
}
