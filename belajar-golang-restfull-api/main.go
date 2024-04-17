package main

import (
	"belajar-golang-restfull-api/src/config"
	"belajar-golang-restfull-api/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := config.DB()

	routes.Api(r, db)

	r.Run()
}
