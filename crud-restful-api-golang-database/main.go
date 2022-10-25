package main

import (
	"crud-restful-api-golang-database/Config"
	"crud-restful-api-golang-database/Models"
	"crud-restful-api-golang-database/Routes"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Book{})

	r := Routes.SetupRouter()
	//running
	r.Run()
}
