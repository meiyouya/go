package main

import (
	"go_study/go_xorm/xorm_engine"
)

func main() {
	xorm_engine.InitConfig()

	xorm_engine.CreateEngine()
}
