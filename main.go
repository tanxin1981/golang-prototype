package main

import (
	_ "golang-prototype/docs"
	_ "golang-prototype/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// register database
	orm.RegisterDataBase("default", "mysql", "root@/sgfas?charset=utf8", 30)
}

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
