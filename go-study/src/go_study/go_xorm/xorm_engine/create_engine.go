package xorm_engine

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"os"
)

func CreateEngine() {
	var engine *xorm.Engine

	mysqlHost := CONFIG.Section("mysql").Key("mysql_host")
	mysqlPort := CONFIG.Section("mysql").Key("mysql_port")
	mysqlUsername := CONFIG.Section("mysql").Key("mysql_username")
	mysqlPassword := CONFIG.Section("mysql").Key("mysql_password")
	mysqlDbName := CONFIG.Section("mysql").Key("mysql_db_name")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", mysqlUsername, mysqlPassword, mysqlHost, mysqlPort, mysqlDbName)
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

	tablePrefixMapper := core.NewPrefixMapper(core.SnakeMapper{}, "lawliet_") // 设置前缀
	engine.SetTableMapper(tablePrefixMapper)                                  // 设置结构体对应的表明前缀

	engine.ShowSQL(true)                     // 显示SQL
	engine.Logger().SetLevel(core.LOG_DEBUG) // 设置日志级别

	file, err := os.Create("go-study/sql.log") // 创建一个用于存储日志的文件
	if err != nil {                            // 创建出错
		println(err.Error())
		return
	}
	engine.SetLogger(xorm.NewSimpleLogger(file)) // 将日志存储到文件中
}
