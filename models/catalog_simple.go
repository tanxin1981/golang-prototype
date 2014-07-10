package models

import (
    "github.com/astaxie/beego/orm"
    "errors"
)

type CatalogSimple struct {
	Id              int               `orm:"column(id_catalog_simple);auto"`
	SizeBrand       string            `orm:"column(size_brand);size(255)"`
    CatalogConfig   *CatalogConfig    `orm:"column(fk_catalog_config);rel(fk)"`
}

func init() {
    // register model
    orm.RegisterModel(new(CatalogSimple))
}

func GetCatalogSimple(simpleId int) (catalogSimple *CatalogSimple, err error) {
    // create a orm instance and query db
    o := orm.NewOrm()
    //orm.Debug = true

    simple := &CatalogSimple{Id: simpleId}
    //err = o.QueryTable("catalog_simple").Filter("CatalogConfig", simpleId).RelatedSel().One(simple)
    err = o.Read(simple)
    if err == nil {
        return simple, nil
    }
    return nil, errors.New("CatalogSimple Not Exist")
}