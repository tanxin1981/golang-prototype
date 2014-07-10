package controllers

import (
	"golang-prototype/models"

	"github.com/astaxie/beego"
	"strconv"
)

// Operations about object
type CatalogController struct {
	beego.Controller
}

func (this *CatalogController) Get() {
	simpleId := this.Ctx.Input.Params[":simpleId"]
	if simpleId != "" {
		id, _ := strconv.Atoi(simpleId)
		catalogSimple, err := models.GetCatalogSimple(int(id))
		if err != nil {
			this.Data["json"] = err
		} else {
			this.Data["json"] = catalogSimple
		}
	}
	this.ServeJson()
}