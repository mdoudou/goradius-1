package goradius

import (
	"encoding/json"
	"io/ioutil"
)

//获取数据库配置
func GetDbConfig() (t, m string) {
	data, err := ioutil.ReadFile("./conf.json")
	if err != nil {
		return
	}
	sqlConfig := &DbConfig{}
	err = json.Unmarshal(data, sqlConfig)
	if err != nil {
		return
	}
	return sqlConfig.SqlType, sqlConfig.SqlCmd
}
