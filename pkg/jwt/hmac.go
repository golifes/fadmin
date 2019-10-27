package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"strings"
)

func Hmac(keys []string, body []byte, sign string) bool {
	s := strings.Join(keys, "_")
	hash := hmac.New(sha256.New, []byte(s))
	hash.Write(body)
	sum := hash.Sum(nil)
	_sign := fmt.Sprintf("%x\n", sum)
	fmt.Println(_sign)
	if fmt.Sprintf("%x\n", sum) == sign {
		return true
	}
	return false

}
