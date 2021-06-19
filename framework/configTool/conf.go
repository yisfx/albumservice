package configTool

import (
	"encoding/json"

	"albumservice/framework/fileTool"
	"albumservice/framework/model"
)

// ReadSysConf
// 读取sysConfig
func ReadSysConf() *model.SysConf {
	conf := &model.SysConf{}
	f := fileTool.GetFileContentByName("./conf/sys.conf.json")
	json.Unmarshal([]byte(f), conf)
	return conf
}

func ReadGlobalConf(path string) *model.GlobalConf {
	conf := &model.GlobalConf{}
	f := fileTool.GetFileContentByName(path)
	json.Unmarshal([]byte(f), conf)
	return conf
}