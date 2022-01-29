package main

import (
	"github.com/gin-gonic/gin"
	"jsonresume/configs"
	"jsonresume/routes"
	"os"
)

func main() {
	os.Setenv("MONGO_URI", "mongodb+srv://kevluxdev:sTTdvsXKDzaJDrgFRxA7dCF4gCT83zBN69RmLWy8Xh9QtkjmCPQjTzpCXWjkfs772VcQCh@resume.quqep.mongodb.net/resume?retryWrites=true&w=majority")
	r := gin.Default()
	r.Static("/css", "./static/css")
	r.LoadHTMLGlob("templates/*.html")
	configs.ConnectDB()
	routes.ResumeRoute(r)
	r.Run("localhost:8000")

}
