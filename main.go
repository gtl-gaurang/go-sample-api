package main

import (
	"fmt"
	Config "simple-api/config"
	Models "simple-api/models"
	Routes "simple-api/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Users{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
