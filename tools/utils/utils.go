package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"regexp"
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
		pn = 0
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

func StringJoin(a ...string) string {
	var buf bytes.Buffer
	for _, k := range a {
		buf.WriteString(k)
	}
	return buf.String()
}

func Slice(query []string, values []interface{}, key string, value interface{}) ([]string, []interface{}) {
	if query != nil {
		query = append(query, fmt.Sprintf(" and %s", key))
		values = append(values, value)
	} else {
		query = append(query, key)
		values = append(values, value)
	}
	return query, values
}

var xss = regexp.MustCompile("[%--`~!@#$^&*()=|{}':;<>/?]")

func Xss(m map[string]interface{}) bool {
	for _, v := range m {
		switch v.(type) {
		case string:
			return xss.MatchString(v.(string))
		default:
			return false
		}
	}
	return false
}
