package key

import (
	"math/rand"
	"strconv"
	"strings"
)

type KeyUint64 uint64

func (u KeyUint64) Equal(u2 KeyUint64) bool {
	return u.ToUint64() == u2.ToUint64()
}

func (u *KeyUint64) FromString(str string) bool {
	value, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	*u = KeyUint64(value)
	return true
}

func (u KeyUint64) String() string {
	return strconv.FormatUint(uint64(u), 10)
}
func (u KeyUint64) ToString() string {
	return strconv.FormatUint(uint64(u), 10)
}
func (u KeyUint64) ToUint64() uint64 {
	return uint64(u)
}

func (u KeyUint64) ToInt() int {
	return int(u)
}
func (u KeyUint64) ToSum() int {
	return int(u)
}
func (u KeyUint64) SumLen() int {
	return 1
}
func (u KeyUint64) GetSumByIdx(_ int) Sum {
	return u
}

type KeyUint64List []KeyUint64

func (keys KeyUint64List) SumLen() int {
	return len(keys)
}
func (keys KeyUint64List) ToString() string {
	return keys.Join("|")
}
func (keys KeyUint64List) ToSum() (sum int) {
	for _, k := range keys {
		sum += k.ToSum()
	}
	return
}
func (keys KeyUint64List) GetSumByIdx(idx int) Sum {
	if idx < len(keys) && idx > -1 {
		return keys[idx]
	}
	return nil
}

func (keys KeyUint64List) ToStringList() (strList []string) {
	strList = make([]string, len(keys))
	for idx, _ := range keys {
		strList[idx] = keys[idx].ToString()
	}
	return
}
func (keys KeyUint64List) Join(sep string) string {
	return strings.Join(keys.ToStringList(), sep)
}

func (keys KeyUint64List) IsInList(Key KeyUint64) bool {
	for _, tmpKey := range keys {
		if tmpKey.Equal(Key) {
			return true
		}
	}
	return false
}

func (keys KeyUint64List) Shuffle() {
	for i := range keys {
		j := rand.Intn(i + 1)
		keys[i], keys[j] = keys[j], keys[i]
	}
}
