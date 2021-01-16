package framework

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../model"
)

func ReadSysConf() *model.SysConf {
	conf := &model.SysConf{}
	f, err := ioutil.ReadFile("./conf/sys.conf.json")
	if err != nil {
		fmt.Println("sys.conf.json error")
		return nil
	}
	json.Unmarshal(f, conf)
	return conf
}

func ReadConf() {

}
