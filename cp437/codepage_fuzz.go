// +build gofuzz

package cp437

import (
	"bytes"
)

func Fuzz(data []byte) int {
	s, err := Decode(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	bs, err := EncodeString(s)
	if err != nil || bytes.Compare(data, bs) != 0 {
		return 0
	}

	return 1
}
