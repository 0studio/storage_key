package key

import (
	"strconv"
	"strings"
)

type String string

func (this String) ToString() string {
	return string(this)
}
func (this String) ToInt() (value int) {
	var err error
	value, err = strconv.Atoi(this.ToString())
	if err == nil {
		return
	}
	return
}

func (this String) ToSum() (value int) {
	for _, byteValue := range this.ToString() {
		value += int(byteValue)
	}
	return
}
func (u String) Equal(u2 String) bool {
	return u == u2
}

type StringList []String

func (keys StringList) ToStringList() (strList []string) {
	strList = make([]string, len(keys))
	for idx, _ := range keys {
		strList[idx] = keys[idx].ToString()
	}
	return
}

func (keys StringList) Join(sep string) string {
	return strings.Join(keys.ToStringList(), sep)
}

func (keys StringList) IsInList(Key String) bool {
	for _, tmpKey := range keys {
		if tmpKey.Equal(Key) {
			return true
		}
	}
	return false
}

type Strings struct {
	key1     Key
	key2     Key
	addition uint32 // 写数据库时用作标示insert or update ， delete时 用作双主键删除的第二个主键的位置
}

func CreateStrings(addition uint32, keys ...Key) Strings {
	strings := Strings{addition: addition}
	if len(keys) >= 1 {
		strings.key1 = keys[0]
	}
	if len(keys) >= 2 {
		strings.key2 = keys[1]
	}
	return strings
}

func (this Strings) ToString() string {
	var stringKey string
	if this.key1 != nil {
		stringKey += this.key1.ToString()
	}
	if this.key2 != nil {
		stringKey += ","
		stringKey += this.key2.ToString()
	}
	return stringKey
}

func (this Strings) Addition() uint32 {
	return this.addition
}

func (this Strings) KeyOne() Key {
	return this.key1
}

func (this Strings) Key(n int) Key {
	if n == 0 {
		return this.key1
	} else if n == 1 {
		return this.key2
	}
	return String("")
}

func (this Strings) KeyCnt() (n int32) {
	if this.key1 != nil {
		n++
	}
	if this.key2 != nil {
		n++
	}
	return
}
