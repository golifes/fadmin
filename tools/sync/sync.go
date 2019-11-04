package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"strings"
)

func main() {

	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"abc123456",
		"49.232.48.41",
		"fadmin")
	fmt.Println(dns)
	EngDb, err := xorm.NewEngine("mysql", dns)

	ping := EngDb.Ping()
	if ping != nil || err != nil {
		panic(ping)
	}

	//EngDb.Sync2(new(wx.WeiXin))
	//EngDb.Sync2(new(admin.DomainAppRole))
	//EngDb.Sync2(new(admin.DomainAppUser))
	//EngDb.Sync2(new(wx.WeiXinList))
	//EngDb.Sync2(new(wx.WeiXin))
	//EngDb.Sync2(new(admin.Group))
	//EngDb.Sync2(new(admin.Role))
	//EngDb.Sync2(new(admin.User))
	//EngDb.Sync2(new(admin.UserRole))
	//EngDb.Sync2(new(admin.UserGroup))
	a := "http://mp.weixin.qq.com/s?__biz=MzU3ODE2NTMxNQ==&MID=2247485961&idx=1&sn=431af867d04efd973fd16df359365dd6&chksm=fd78c525ca0f4c334da2c677c1622f32058b7d3b89d255d5bb6e21a11a7f32407b67b13245bd&scene=27#wechat_redirect"

	arr := FindBizStr(a)
	fmt.Println(arr)
}

func FindBizStr(url string) (arr []string) {
	fmt.Println(url)

	//index := strings.Index(a, "__biz=")
	//fmt.Println(index)
	//print(a[index:])
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
