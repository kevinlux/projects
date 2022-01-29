package main

import (
	"github.com/gin-gonic/gin"
	"jsonresume/configs"
	"jsonresume/routes"
	"os"
)

func main() {
	r := gin.Default()
	r.Static("/css", "./static/css")
	r.LoadHTMLGlob("templates/*.html")
	configs.ConnectDB()
	routes.ResumeRoute(r)
	r.Run("localhost:8000")

}
