package xorm_engine

import (
	"github.com/go-ini/ini"
	"os"
)

var CONFIG *ini.File

/**
初始化配置
*/
func InitConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		println(err.Error())
		os.Exit(2)
		return
	}
	CONFIG = cfg
}
