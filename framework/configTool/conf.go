package configTool

import (
	"encoding/json"

	"albumservice/framework/bootstrapmodel"
	"albumservice/framework/fileTool"
	"albumservice/framework/utils"
)

// ReadSysConf
// 读取sysConfig
func ReadSysConf() *bootstrapmodel.SysConf {
	conf := &bootstrapmodel.SysConf{}
	f := fileTool.GetFileContentByName("./conf/sys.conf.json")
	json.Unmarshal([]byte(f), conf)
	return conf
}

func ReadGlobalConf(path string) *bootstrapmodel.GlobalConf {
	defer utils.HanderError()

	conf := &bootstrapmodel.GlobalConf{}
	f := fileTool.GetFileContentByName(path)
	json.Unmarshal([]byte(f), conf)
	return conf
}
