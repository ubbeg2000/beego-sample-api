package main

import (
	database "sample-api/database"
	"sample-api/models"
	_ "sample-api/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	database.Conn.AutoMigrate(
		&models.Todo{},
		&models.User{},
	)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
