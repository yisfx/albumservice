package framework

import (
	"encoding/json"

	model "albumservice/model"
)

// ReadSysConf
// 读取sysConfig
func ReadSysConf() *model.SysConf {
	conf := &model.SysConf{}
	f := GetFileContentByName("./conf/sys.conf.json")
	json.Unmarshal([]byte(f), conf)
	return conf
}

func ReadGlobalConf(path string) *model.GlobalConf {
	conf := &model.GlobalConf{}
	f := GetFileContentByName(path)
	json.Unmarshal([]byte(f), conf)
	return conf
}
