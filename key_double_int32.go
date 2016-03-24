package key

import (
	"fmt"
	"strconv"
)

type KeyDoubleInt32 struct {
	key1 int32
	key2 int32
}

func NewKeyDoubleInt32(key1, key2 int32) KeyDoubleInt32 {
	return KeyDoubleInt32{key1: key1, key2: key2}
}
func (u KeyDoubleInt32) Equal(u2 KeyDoubleInt32) bool {
	return u.key1 == u2.key1 && u.key2 == u2.key2
}

func (u *KeyDoubleInt32) FromString(str string) bool {
	_, err := fmt.Sscanf(str, "%d-%d", &u.key1, &u.key2)
	return err == nil
}

func (u KeyDoubleInt32) String() string {
	return u.ToString()
}
func (u KeyDoubleInt32) ToString() string {
	return strconv.FormatInt(int64(u.key1), 10) + "-" + strconv.FormatInt(int64(u.key2), 10)
}
func (u KeyDoubleInt32) GetKeys() (key1, key2 int32) {
	key1 = u.key1
	key2 = u.key2
	return
}
func (u KeyDoubleInt32) ToInt() int {
	return int(u.key1 + u.key2)
}
func (u KeyDoubleInt32) SumLen() int {
	return 1
}
func (keys KeyDoubleInt32) ToSum() (sum int) {
	return int(keys.key1 + keys.key2)
}

func (u KeyDoubleInt32) GetSumByIdx(_ int) Sum {
	return u
}

func (u KeyDoubleInt32) GetKey1() (key1 int32) {
	key1 = u.key1
	return
}
func (u KeyDoubleInt32) GetKey2() (key2 int32) {
	key2 = u.key2
	return
}
func (u KeyDoubleInt32) GetOtherKey(myUin int32) (ret int32, ok bool) {
	if myUin == u.key1 {
		ok = true
		ret = u.key2
		return
	} else if myUin == u.key2 {
		ok = true
		ret = u.key1
		return
	}

	return
}
func (u *KeyDoubleInt32) SetKeys(key1, key2 int32) {
	u.key1 = key1
	u.key2 = key2
}

type KeyDoubleInt32List []KeyDoubleInt32

func (keys KeyDoubleInt32List) GetOtherKeyList(myUin int32) (retList []int32) {
	var otherKey int32
	var ok bool
	for _, u := range keys {
		otherKey, ok = u.GetOtherKey(myUin)
		if ok {
			retList = append(retList, otherKey)
		}
	}
	return
}
func (keys KeyDoubleInt32List) Remove(Key KeyDoubleInt32) (ret KeyDoubleInt32List) {
	ret = make(KeyDoubleInt32List, len(keys), len(keys))
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

func (keys KeyDoubleInt32List) IsInList(Key KeyDoubleInt32) bool {
	for _, tmpKey := range keys {
		if tmpKey.Equal(Key) {
			return true
		}
	}
	return false
}
func (u KeyDoubleInt32List) ToInt() int {
	return 0
}
func (u KeyDoubleInt32List) SumLen() int {
	return len(u)
}
func (u KeyDoubleInt32List) GetSumByIdx(idx int) Sum {
	return u[idx]
}
