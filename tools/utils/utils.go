package utils

import (
	"crypto/md5"
	"encoding/hex"
	"log"
)

func CheckError(err error, v interface{}) bool {
	if err != nil {
		log.Printf("err is %s,%s", err, v)
		return false
	}
	return true
}

func Pagination(pageSize, pageNum, defaultPageSize int) (ps, pn int) {

	if pageNum <= 1 {
		pn = 1
	} else {
		pn = pageNum
	}

	if pageSize <= 0 || pageSize >= defaultPageSize {
		ps = defaultPageSize
	} else {
		ps = pageSize
	}
	return
}

func EncodeMd5(value string) string {

	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
