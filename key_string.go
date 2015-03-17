package key

import (
	"strings"
)

type String string

func (this String) ToString() string {
	return string(this)
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
