package gen

import (
	"errors"
	"math/rand"
)

const (
	CharsetDigit     = "0123456789"
	CharsetLowerCase = "abcdefghijklmnopqrstuvwxyz"
	CharsetUpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharsetSpecial   = " ~!@#$%^&*()_+-=[]{};'\\:\"|,./<>?"
)

// StrGenerate
// generate a random string, all chars in this string are in your charset
// and length is in [minLen, maxLen]
func StrGenerate(charset string, minLen, maxLen int) (str string, err error) {
	var bstr []byte
	bstr, err = BytesGenerate([]byte(charset), minLen, maxLen)
	if err != nil {
		return
	}
	str = string(bstr)
	return
}

// BytesGenerate
// generate a random bytes, all byte in this bytes are in your charset
// and length is in [minLen, maxLen]
func BytesGenerate(charset []byte, minLen, maxLen int) (bresult []byte, err error) {
	if minLen > maxLen {
		err = errors.New("minLen must less or equal to maxLen")
		return
	}
	charsetSize := len(charset)
	if charsetSize <= 0 {
		err = errors.New("len(charset) must greater than 0")
		return
	}
	length := rand.Intn(maxLen-minLen+1) + minLen
	bresult = make([]byte, 0, length)
	for i := 0; i < length; i++ {
		b := charset[rand.Intn(charsetSize)]
		bresult = append(bresult, b)
	}
	return
}
