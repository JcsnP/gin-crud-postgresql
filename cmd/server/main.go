package main

import (
	"github.com/JcsnP/gin-clean/config/db"
	"github.com/JcsnP/gin-clean/pkg/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	db := db.ConnectDatabase()
	app := gin.Default()

	routers.RouterSetup(app, db)

	app.Run()
}