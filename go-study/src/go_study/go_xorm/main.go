package main

import "go_study/go_xorm/xorm_engine"

func main() {
	xorm_engine.InitConfig()   // 初始化配置文件
	xorm_engine.CreateEngine() // 创建引擎
}
