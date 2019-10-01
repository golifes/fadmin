package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func Load(path string, v interface{}) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("load读取配置文件出错--->", path)
		panic(err)
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		fmt.Println("load序列化配置文件出错--->", path)
		panic(err)
	}
}
