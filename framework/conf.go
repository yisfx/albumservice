package framework

import (
	"encoding/json"

	"../helper"
	"../model"
)

// ReadSysConf
// 读取sysConfig
func ReadSysConf() *model.SysConf {
	conf := &model.SysConf{}
	f := helper.GetFileContentByName("./conf/sys.conf.json")
	json.Unmarshal([]byte(f), conf)
	return conf
}
