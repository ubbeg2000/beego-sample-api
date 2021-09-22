package database

import (
	"github.com/beego/beego/v2/core/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func init() {
	if rm, err := config.String("runmode"); err == nil {
		if rm == "dev" {
			db, err := gorm.Open(sqlite.Open("dev.db"), &gorm.Config{})

			if err != nil {
				panic(err.Error())
			}

			Conn = db
		}
	} else {
		panic(err.Error())
	}
}
