package jwt

import "testing"

func TestHmac(t *testing.T) {
	Hmac([]string{"abc", "GET", "/api/v1", "2019-10-25 :16:40:23"}, []byte("12345"), "111")
}
