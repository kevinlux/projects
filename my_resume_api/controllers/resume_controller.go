package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"jsonresume/configs"
	"jsonresume/models"
	"jsonresume/responses"
	"net/http"
	"time"
)

var resumeCollection *mongo.Collection = configs.GetCollection(configs.DB, "resume")
var validate = validator.New()

func ResumeFromDB(c *gin.Context) models.Resume {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var resume models.Resume
	defer cancel()

	if err := resumeCollection.FindOne(ctx, bson.M{}).Decode(&resume); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ResumeResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
	}
	return resume

}

func GetHTMLResume() gin.HandlerFunc {
	return func(c *gin.Context) {
		resume := ResumeFromDB(c)
		c.HTML(http.StatusOK, "resume.html", resume)
	}
}

func GetResource(resource string) gin.HandlerFunc {
	return func(c *gin.Context) {
		resume := ResumeFromDB(c)
		switch resource {
		case "resume":
			c.JSON(http.StatusOK, responses.ResumeResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": resume}})
		case "fullname":
			c.JSON(http.StatusOK, responses.ResumeResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": resume.FullName}})
		case "linkedin":
			c.JSON(http.StatusOK, responses.ResumeResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": resume.Linkedin}})
		case "github":
			c.JSON(http.StatusOK, responses.ResumeResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": resume.Github}})
		case "website":
			c.JSON(http.StatusOK, responses.ResumeResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": resume.Website}})
		case "summary":
			c.JSON(http.StatusOK, responses.ResumeResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": resume.Summary}})
		case "technologies":
			c.JSON(http.StatusOK, responses.ResumeResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": resume.Technologies}})
		case "education":
			c.JSON(http.StatusOK, responses.ResumeResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": resume.Education}})
		case "workhistory":
			c.JSON(http.StatusOK, responses.ResumeResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": resume.WorkHistory}})
		case "languages":
			c.JSON(http.StatusOK, responses.ResumeResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": resume.Languages}})
		}
	}
}

func UpdateResume() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.GetHeader("Authorization")
		if bearer != "Bearer "+configs.Bearer {
			c.JSON(http.StatusBadRequest, responses.ResumeResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "invalid bearer token"}})
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var resume models.Resume
		defer cancel()

		if err := c.BindJSON(&resume); err != nil {
			c.JSON(http.StatusBadRequest, responses.ResumeResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		}

		if validationErr := validate.Struct(&resume); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ResumeResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		}

		update := bson.M{
			"fullname":     resume.FullName,
			"linkedin":     resume.Linkedin,
			"github":       resume.Github,
			"website":      resume.Website,
			"summary":      resume.Summary,
			"technologies": resume.Technologies,
			"education":    resume.Education,
			"workhistory":  resume.WorkHistory,
			"languages":    resume.Languages,
		}

		result, err := resumeCollection.UpdateOne(ctx, bson.M{"fullname": "Kevin Lux"}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ResumeResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var updatedResume models.Resume
		if result.MatchedCount == 1 {
			err := resumeCollection.FindOne(ctx, bson.M{"fullname": "Kevin Lux"}).Decode(&updatedResume)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.ResumeResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.ResumeResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedResume}})

	}
}
