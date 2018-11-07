package controllers

import (
	//	"errors"
	//	"fmt"

	"consumer/infrastructure/db"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/go-xorm/xorm"
)

// Operations about tag
type TagController struct {
	beego.Controller
}

var engine *xorm.Engine

type info struct {
	name string
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (t *TagController) Get() {

	key := t.GetString("Index")
	value := t.GetString(key)

	if key != "" && value != "" {
		n := db.GetBatchCountViaToken(1)
		logs.Error("dsdsd;%d", n)
		err, name := db.GetName()

		if err != nil {
			t.Data["json"] = err.Error()
		} else {
			t.Data["json"] = name
		}
	}
	t.ServeJSON()
}
