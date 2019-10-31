package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"regexp"
	"strings"
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

func FindBizStr(url string) (arr []string) {
	bizIndex := strings.Index(url, "__biz=")
	if bizIndex == -1 {
		return nil
	}
	bizEnd := strings.Index(url[bizIndex:], "&")
	biz := url[bizIndex+6 : bizEnd+bizIndex]
	arr = append(arr, biz)

	//mid
	midIndex := strings.Index(url, "mid=")
	if midIndex == -1 {
		midIndex = strings.Index(url, "MID=")
	}
	if midIndex == -1 {
		return nil
	}
	midEnd := strings.Index(url[midIndex:], "&")
	mid := url[midIndex+4 : midIndex+midEnd]
	arr = append(arr, mid)

	idxIndex := strings.Index(url, "&idx=")
	if idxIndex == -1 {
		idxIndex = strings.Index(url, "&IDX=")
	}
	if midIndex == -1 {
		return nil
	}
	idxEnd := strings.Index(url[idxIndex+5:], "&")
	idx := url[idxIndex+5 : idxEnd+idxIndex+5]
	arr = append(arr, idx)

	return
}
