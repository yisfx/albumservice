package framework

import (
	"encoding/json"

	"../model"
)

// ReadSysConf
// 读取sysConfig
func ReadSysConf() *model.SysConf {
	conf := &model.SysConf{}
	f := GetFileContentByName("./conf/sys.conf.json")
	json.Unmarshal([]byte(f), conf)
	return conf
}
