package main

import (
	"example.com/mod/config"
	"example.com/mod/models"
	"example.com/mod/router"
)

func main() {
	config.Config()
	router.InitRouter()
	models.InitMysqlDB()
}
