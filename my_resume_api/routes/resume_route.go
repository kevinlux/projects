package routes

import (
	"github.com/gin-gonic/gin"
	"jsonresume/controllers"
)

func ResumeRoute(router *gin.Engine) {
	router.GET("/resume.html", controllers.GetHTMLResume())
	router.GET("/", controllers.GetResource("resume"))
	router.GET("/fullname", controllers.GetResource("fullname"))
	router.GET("/email", controllers.GetResource("email"))
	router.GET("/website", controllers.GetResource("website"))
	router.GET("/summary", controllers.GetResource("summary"))
	router.GET("/technologies", controllers.GetResource("technologies"))
	router.GET("/education", controllers.GetResource("education"))
	router.GET("/workhistory", controllers.GetResource("workhistory"))
	router.GET("/languages", controllers.GetResource("languages"))
	router.PUT("/", controllers.UpdateResume())
}
