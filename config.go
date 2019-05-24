package goradius

import (
	"encoding/json"
	"io/ioutil"
)

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
