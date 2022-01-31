package routes

import (
	"github.com/gin-gonic/gin"
	"jsonresume/controllers"
)

func ResumeRoute(router *gin.Engine) {
	router.GET("/resume.html", controllers.GetHTMLResume())
	router.GET("/", controllers.GetResource("resume"))
	router.GET("/fullname", controllers.GetResource("fullname"))
	router.GET("/linkedin", controllers.GetResource("linkedin"))
	router.GET("/github", controllers.GetResource("github"))
	router.GET("/website", controllers.GetResource("website"))
	router.GET("/summary", controllers.GetResource("summary"))
	router.GET("/technologies", controllers.GetResource("technologies"))
	router.GET("/education", controllers.GetResource("education"))
	router.GET("/workhistory", controllers.GetResource("workhistory"))
	router.GET("/languages", controllers.GetResource("languages"))
	router.PUT("/", controllers.UpdateResume())
}
