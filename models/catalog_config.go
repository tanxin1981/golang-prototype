package models

import (
	"github.com/astaxie/beego/orm"
)

type CatalogConfig struct {
	Id		int		`orm:"auto;column(id_catalog_config)"`
	Name	string	`orm:"column(name)"`
	Color	string	`orm:"column(color)"`
}

func init() {
	// register model
    orm.RegisterModel(new(CatalogConfig))
}