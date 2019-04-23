package xorm_engine

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"os"
)

var engine *xorm.Engine

func CreateEngine() {

	host := CONFIG.Section("mysql").Key("mysql_host")
	port := CONFIG.Section("mysql").Key("mysql_port")
	username := CONFIG.Section("mysql").Key("mysql_username")
	password := CONFIG.Section("mysql").Key("mysql_password")

	dataSourceName := "root:521xwl54zql@MYSQL@tcp(39.108.92.8:3306)/test?charset=utf8"
	engine, err := xorm.NewEngine("mysql", dataSourceName) // 创建连接引擎，url格式   username:password@tcp(host:port)/dbName?otherParams
	if err != nil {
		println(err.Error())
		return
	}

	err = engine.Ping() // PING数据库
	if err != nil {
		println(err.Error())
		return
	}

	engine.ShowSQL(true)                     // 显示SQL
	engine.Logger().SetLevel(core.LOG_DEBUG) // 设置日志级别

	file, err := os.Create("sql.log") // 创建一个用于存储日志的文件
	if err != nil {                   // 创建出错
		println(err.Error())
		return
	}
	engine.SetLogger(xorm.NewSimpleLogger(file)) // 将日志存储到文件中
}
