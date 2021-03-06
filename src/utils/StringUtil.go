package utils

import (
	//"time"
	//"strconv"
	//"os"
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"strings"
)

func GetFromHex(hexSource string) (to []byte, err error) {
	bs, err := hex.DecodeString(hexSource)
	if err != nil {
		Errorf("hex.DecodeString %s", err)
		return nil, err
	}
	return bs, nil
}

//byteSource = []byte("大家好")
func GetFromUtf8(bs []byte) string {
	//enc2 := mahonia.NewEncoder("UTF-8")
	//enc2.ConvertString()
	return string(bs)
}

func GetFromUnicode(source string) (to string, err error) {
	bs, err := hex.DecodeString(strings.Replace(source, `\u`, ``, -1))
	if err != nil {
		return
	}
	for i, bl, br, r := 0, len(bs), bytes.NewReader(bs), uint16(0); i < bl; i += 2 {
		binary.Read(br, binary.BigEndian, &r)
		to += string(r)
	}
	return
}
