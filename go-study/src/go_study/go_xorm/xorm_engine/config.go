package xorm_engine

import (
	"github.com/go-ini/ini"
	"os"
)

var CONFIG *ini.File // 全局变量   CONFIG，存储config.ini中的信息

/**
初始化配置
*/
func InitConfig() {
	cfg, err := ini.Load("go-study/src/go_study/go_xorm/config.ini")
	if err != nil {
		println(err.Error())
		os.Exit(2)
		return
	}
	CONFIG = cfg
}
