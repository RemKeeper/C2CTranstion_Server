package SettingDbCtrl

import (
	"C2CTranstion_Server/DbInit"
	"log"
	"os"
)

var (
	Db           = DbInit.Db
	GlobalConfig Config
)

func init() {
	var err error
	GlobalConfig, err = GetSetting()
	if err != nil {
		log.Println("获取全局配置失败")
		os.Exit(404)
		return
	}
}
