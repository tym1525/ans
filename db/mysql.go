package db

import (
	"answersys/libs/module"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	InitMysql("anssys:tym1525@tcp(127.0.0.1:3306)/anssys?charset=utf8&loc=Local", 30)
}

func InitMysql(con string, max int) {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", con, max)
	orm.RegisterModelWithPrefix("ans_", new(module.User))
	orm.RunSyncdb("default", false, true)
}
