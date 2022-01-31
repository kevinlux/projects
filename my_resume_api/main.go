package main

import (
	"github.com/gin-gonic/gin"
	"jsonresume/configs"
	"jsonresume/routes"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	r := gin.New()
	r.Use(gin.Logger())
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/css", "./static/css")
	configs.ConnectDB()
	routes.ResumeRoute(r)
	r.Run(":" + port)
}
