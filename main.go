package main

import (
	_ "consumer/infrastructure/db"
	_ "consumer/routers"

	"consumer/models"

	"github.com/astaxie/beego"
)

func main() {
	//db.init()
	if beego.BConfig.RunMode == "prod" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	iskafka, _ := beego.AppConfig.Bool("kafka::mode")
	if iskafka {
		go models.Consumer()
	}

	go models.DgraphClient()
	beego.Run()
}
